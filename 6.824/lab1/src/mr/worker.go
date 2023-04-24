package mr

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strings"
)

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// Your worker implementation here.
	lastType, lastID, output := "", 0, []string{}

	id := os.Getpid()
	for {
		args := AskArgs{
			WorkerID:     id,
			LastTaskType: lastType,
			LastTaskID:   lastID,
			Output:       output,
		}
		reply := AskReply{}
		call("Coordinator.Ask", &args, &reply)

		fmt.Println(reply)
		if reply.NewTask.Type == "" {
			log.Printf("worker: %d done", id)
			break
		}
		if reply.NewTask.Type == Map {
			output = mapFunc(reply.NewTask, reply.NReduce, mapf)
		}

		if reply.NewTask.Type == Reduce {
			output = reduceFunc(reply.NewTask, reducef)
		}
		lastType = reply.NewTask.Type
		lastID = reply.NewTask.ID
	}

	// uncomment to send the Example RPC to the coordinator.
	// CallExample()

}

func mapFunc(task Task, nReduce int, mapf func(string, string) []KeyValue) []string {
	if len(task.Input) < 1 {
		return []string{}
	}
	file, err := os.Open(task.Input[0])
	if err != nil {
		log.Fatalf("open file failed, err: %v", err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("read file content failed, err: %v", err)
	}

	sli := mapf(task.Input[0], string(content))
	m := make(map[int][]KeyValue)
	for _, kv := range sli {
		hashed := ihash(kv.Key) % nReduce
		m[hashed] = append(m[hashed], kv)
	}

	res := make([]string, nReduce)
	for i := 0; i < nReduce; i++ {
		p := fmt.Sprintf("map-%d-%d-%d", task.WorkerID, task.ID, i)
		f, err := os.OpenFile(
			p,
			os.O_RDWR|os.O_CREATE|os.O_APPEND,
			0666,
		)
		if err != nil {
			log.Fatal(err)
		}
		for _, kv := range m[i] {
			fmt.Fprintf(f, "%s\t%s\n", kv.Key, kv.Value)
		}
		res[i] = p
		f.Close()
	}

	return res
}

func reduceFunc(task Task, reducef func(string, []string) string) []string {
	lines := []string{}
	for _, p := range task.Input {
		f, err := os.Open(p)
		if err != nil {
			log.Fatalf("open file failed, err: %v", err)
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("read file failed, err: %v", err)
		}
		lines = append(lines, strings.Split(string(content), "\n")...)
	}

	sli := make([]KeyValue, 0, len(lines))
	for _, line := range lines {
		kv := strings.Split(line, "\t")
		if len(kv) < 2 {
			continue
		}
		sli = append(sli, KeyValue{
			Key:   kv[0],
			Value: kv[1],
		})
	}

	sort.Slice(sli, func(i, j int) bool {
		return sli[i].Key < sli[j].Key
	})

	name := fmt.Sprintf("reduce-%d-%d", task.WorkerID, task.ID)
	ofile, err := os.Create(name)
	if err != nil {
		log.Fatalf("create file failed, err: %v", err)
	}

	i := 0
	for i < len(sli) {
		j := i + 1
		for j < len(sli) && sli[i].Key == sli[j].Key {
			j++
		}
		vals := []string{}
		for k := i; k < j; k++ {
			vals = append(vals, sli[k].Value)
		}

		output := reducef(sli[i].Key, vals)

		fmt.Fprintf(ofile, "%s %s\n", sli[i].Key, output)
		i = j
	}
	ofile.Close()

	return []string{name}
}

// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	call("Coordinator.Example", &args, &reply)

	// reply.Y should be 100.
	fmt.Printf("reply.Y %v\n", reply.Y)
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
