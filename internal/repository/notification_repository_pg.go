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

func (n *notificationRepositoryPG) Create(ctx context.Context, task domain.Task) (result domain.Task, err error) {
	err = n.pgSession.
		InsertInto("task").
		Columns("run_time", "email", "name").
		Record(task).
		Returning("id", "run_time", "email", "name").
		LoadContext(ctx, &result)
	return
}

func (n *notificationRepositoryPG) GetTasks(from int64, to int64) (result domain.Tasks, err error) {
	_, err = n.pgSession.
		Select("id", "run_time", "email", "name").
		From("task").
		Where("run_time between ? and ?", from, to).
		LoadContext(context.Background(), &result)
	return
}

func (n *notificationRepositoryPG) Delete(id int64) (err error) {
	_, err = n.pgSession.
		DeleteFrom("task").
		Where("id = ?", id).
		ExecContext(context.Background())
	return
}
