package model

import "context"

type TasksModel interface {
	tasksModelReader
	tasksModelWriter
}

type tasksModelReader interface {
	FindTasks(ctx context.Context) ([]*Task, error)
}

type tasksModelWriter interface{}

type UserDailyTaskModel interface {
	userDailyTaskReader
	userDailyTaskWriter
}

type userDailyTaskReader interface {
}

type userDailyTaskWriter interface {
}
