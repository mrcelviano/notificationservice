package domain

import (
	"context"
)

type Task struct {
	ID      int64  `json:"id" db:"id"`
	RunTime int64  `json:"runTime" db:"run_time"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
}

type NotificationService interface {
	RegisterTask(context.Context, Task) (int64, error)
	StartNotificationScheduler()
}

type NotificationRepositoryPG interface {
	Create(context.Context, Task) (Task, error)
	GetTasks(int64, int64) ([]Task, error)
	Delete(int64) error
}
