package mr

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

const (
	Map    = "Map"
	Reduce = "Reduce"
	Done   = "done"
)

type Coordinator struct {
	// Your definitions here.
	mu        sync.Mutex
	stage     string
	nMap      int
	nReduce   int
	id        int
	tasks     map[int]*Task
	available chan *Task
}

// Your code here -- RPC handlers for the worker to call.

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) Ask(args *AskArgs, reply *AskReply) error {
	if args == nil || reply == nil {
		return errors.New("param error")
	}

	task, ok := c.getTask(args.WorkerID, args.LastTaskType, args.LastTaskID, args.Output)
	if !ok {
		return nil
	}

	reply.NewTask = task
	reply.NMap = c.nMap
	reply.NReduce = c.nReduce

	return nil
}

func (c *Coordinator) getTask(
	worker int,
	lastTaskType string,
	lastTaskID int,
	output []string,
) (Task, bool) {
	switch lastTaskType {
	case "":
		return c.getNewTask(worker, lastTaskID)
	case Map:
		return c.getNewTaskAfterMap(worker, lastTaskID, output)
	case Reduce:
		return c.getNewTaskAfterReduce(worker, lastTaskID, output)
	}

	return Task{}, false
}

func (c *Coordinator) getNewTask(worker, lastTaskID int) (Task, bool) {
	task, ok := <-c.available
	if !ok {
		return Task{}, false
	}
	c.mu.Lock()
	task.Deadline = time.Now().Add(time.Second * 10)
	task.WorkerID = worker
	c.mu.Unlock()

	return *task, true
}

func (c *Coordinator) getNewTaskAfterMap(worker, lastTaskID int, output []string) (Task, bool) {
	c.mu.Lock()
	if c.tasks[lastTaskID].WorkerID == worker {
		delete(c.tasks, lastTaskID)
		for i, f := range output {
			os.Rename(f, c.getMapOutput(lastTaskID, i))

		}
		if len(c.tasks) == 0 {
			c.stage = Reduce
			for i := 0; i < c.nReduce; i++ {
				c.id++
				task := &Task{
					ID:   c.id,
					Type: Reduce,
				}
				for j := 0; j < c.nMap; j++ {
					task.Input = append(task.Input, c.getMapOutput(j, i))
				}
				c.tasks[task.ID] = task
				c.available <- task
			}
		}
	} else {
		for _, f := range output {
			os.Remove(f)
		}
	}
	c.mu.Unlock()
	task, ok := <-c.available
	if !ok {
		return Task{}, false
	}
	c.mu.Lock()
	task.Deadline = time.Now().Add(time.Second * 10)
	task.WorkerID = worker
	c.mu.Unlock()

	return *task, true
}

func (c *Coordinator) getNewTaskAfterReduce(worker, lastTaskID int, output []string) (Task, bool) {
	c.mu.Lock()
	if c.tasks[lastTaskID].WorkerID == worker {
		for _, f := range output {
			os.Rename(f, c.getReduceOutput(lastTaskID))
		}
		for _, f := range c.tasks[lastTaskID].Input {
			os.Remove(f)
		}
		delete(c.tasks, lastTaskID)
		if len(c.tasks) == 0 {
			c.stage = Done
			close(c.available)
		}
	} else {
		for _, f := range output {
			os.Remove(f)
		}
	}
	c.mu.Unlock()
	task, ok := <-c.available
	if !ok {
		return Task{}, false
	}
	c.mu.Lock()
	task.Deadline = time.Now().Add(time.Second * 10)
	task.WorkerID = worker
	c.mu.Unlock()

	return *task, true
}

func (c *Coordinator) getMapOutput(id, r int) string {
	return fmt.Sprintf("map-%d-%d", id, r)
}

func (c *Coordinator) getReduceOutput(id int) string {
	return fmt.Sprintf("mr-out-%d", id)
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

}

func (c *Coordinator) check() {
	for {
		time.Sleep(time.Microsecond * 500)

		c.mu.Lock()
		for _, task := range c.tasks {
			if task.WorkerID != 0 && time.Now().After(task.Deadline) {
				task.WorkerID = 0
				c.available <- task
			}
		}
		c.mu.Unlock()
	}
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.

	return ret
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		stage:     Map,
		nMap:      len(files),
		nReduce:   nReduce,
		tasks:     make(map[int]*Task),
		available: make(chan *Task, max(len(files), nReduce)),
	}

	for i, file := range files {
		task := Task{
			Type:  Map,
			ID:    i,
			Input: []string{file},
		}
		c.available <- &task
		c.tasks[i] = &task
	}

	c.id = len(files) - 1

	c.server()

	c.check()

	return &c
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
