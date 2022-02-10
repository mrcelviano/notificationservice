package app

import (
	"context"
)

type Tasks []Task

type Task struct {
	ID      int64  `json:"id" db:"id"`
	RunTime int64  `json:"runTime" db:"run_time"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
}

type NotificationLogic interface {
	RegisterTask(context.Context, Task) (int64, error)
	Start()
}

type NotificationRepository interface {
	Create(context.Context, Task) (Task, error)
	GetTasks(int64, int64) (Tasks, error)
	Delete(int64) error
}

type SendLogic interface {
	SendNotification(string, string)
}
