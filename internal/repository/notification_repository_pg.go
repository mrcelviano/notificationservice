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

func (n *notificationRepositoryPG) Create(ctx context.Context, task domain.Task) error {
	_, err := n.pgSession.
		InsertInto("task").
		Columns("run_time", "user_id").
		Record(task).
		ExecContext(ctx)
	if err != nil {
		return domain.ErrCantExecSQLRequest
	}
	return nil
}

func (n *notificationRepositoryPG) GetTasks(from int64, to int64) ([]domain.Task, error) {
	var taskList []domain.Task
	_, err := n.pgSession.
		Select("id", "run_time", "user_id").
		From("task").
		Where("run_time between ? and ?", from, to).
		LoadContext(context.Background(), &taskList)
	if err != nil {
		return nil, domain.ErrCantExecSQLRequest
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
