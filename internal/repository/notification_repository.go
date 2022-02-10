package repository

import (
	"context"
	"github.com/gocraft/dbr"
	"github.com/mrcelviano/notificationservice/commons"
	"github.com/mrcelviano/notificationservice/internal/app"
)

type notificationRepository struct {
	tagMap     map[string]string
	retCols    []string
	createCols []string
	pgConn     *dbr.Connection
}

func NewNotificationRepository(pgConn *dbr.Connection) app.NotificationRepository {
	return &notificationRepository{
		tagMap:     commons.MakeJsonToDbTagMap(app.Task{}),
		retCols:    commons.FindDbTags(app.Task{}),
		createCols: commons.FindDbTags(app.Task{}, "id"),
		pgConn:     pgConn,
	}
}

func (n *notificationRepository) Create(ctx context.Context, task app.Task) (result app.Task, err error) {
	sess := commons.DBSessionFromContext(ctx)

	err = sess.
		InsertInto("task").
		Columns(n.createCols...).
		Record(task).
		Returning(n.retCols...).
		LoadContext(ctx, &result)
	return
}

func (n *notificationRepository) GetTasks(from int64, to int64) (result app.Tasks, err error) {
	ctx := commons.DBSessionNewContext(context.Background(), n.pgConn)
	sess := commons.DBSessionFromContext(ctx)

	_, err = sess.
		Select(n.retCols...).
		From("task").
		Where("run_time between ? and ?", from, to).
		LoadContext(ctx, &result)
	return
}

func (n *notificationRepository) Delete(id int64) (err error) {
	ctx := commons.DBSessionNewContext(context.Background(), n.pgConn)
	sess := commons.DBSessionFromContext(ctx)

	_, err = sess.
		DeleteFrom("task").
		Where("id = ?", id).
		ExecContext(ctx)
	return
}
