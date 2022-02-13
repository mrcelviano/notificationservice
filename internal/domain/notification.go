package domain

import (
	"context"
)

type Task struct {
	ID      int64 `json:"id"`
	RunTime int64 `json:"runTime"`
	UserID  int64 `json:"userID"`
}

type NotificationService interface {
	RegisterTask(context.Context, int64) (bool, error)
	StartNotificationScheduler()
}

type NotificationRepositoryPG interface {
	Create(context.Context, Task) error
	GetTasks(int64, int64) ([]Task, error)
	Delete(int64) error
}
