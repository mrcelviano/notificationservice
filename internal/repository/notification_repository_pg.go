package repository

import (
	"context"
	goCraft "github.com/gocraft/dbr"
	"github.com/mrcelviano/notificationservice/internal/domain"
)

type notificationRepositoryPG struct {
	pgSession *goCraft.Session
}

func NewNotificationRepositoryPG(pgDBConn *goCraft.Connection) domain.NotificationRepositoryPG {
	return &notificationRepositoryPG{
		pgSession: pgDBConn.NewSession(&eventReceiver{}),
	}
}

func (n *notificationRepositoryPG) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	var newTask domain.Task
	err := n.pgSession.
		InsertInto("task").
		Columns("run_time", "email", "name").
		Record(task).
		Returning("id", "run_time", "email", "name").
		LoadContext(ctx, &newTask)
	if err != nil {
		return newTask, err
	}
	return newTask, nil
}

func (n *notificationRepositoryPG) GetTasks(from int64, to int64) ([]domain.Task, error) {
	var taskList []domain.Task
	_, err := n.pgSession.
		Select("id", "run_time", "email", "name").
		From("task").
		Where("run_time between ? and ?", from, to).
		LoadContext(context.Background(), &taskList)
	if err != nil {
		return nil, err
	}
	return taskList, nil
}

func (n *notificationRepositoryPG) Delete(id int64) error {
	_, err := n.pgSession.
		DeleteFrom("task").
		Where("id = ?", id).
		ExecContext(context.Background())
	if err != nil {
		return err
	}
	return nil
}
