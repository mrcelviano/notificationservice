package service

import (
	"context"
	"github.com/mrcelviano/notificationservice/internal/domain"
	"github.com/mrcelviano/notificationservice/pkg/logger"
	"github.com/mrcelviano/notificationservice/pkg/scheduler"
	"time"
)

const (
	timeValue       = time.Minute * 30
	countGoroutines = 5
)

type notificationService struct {
	repo domain.NotificationRepositoryPG
	send domain.SendService

	chTasks  chan domain.Task
	fromTime int64
	toTime   int64
}

func NewNotificationService(repo domain.NotificationRepositoryPG, send domain.SendService) domain.NotificationService {
	return &notificationService{
		repo: repo,
		send: send,
	}
}

func (n *notificationService) RegisterTask(ctx context.Context, task domain.Task) (id int64, err error) {
	task.RunTime = time.Now().Add(timeValue).Unix()
	task, err = n.repo.Create(ctx, task)
	if err != nil {
		return id, domain.ErrCantExecSQLRequest
	}
	return task.ID, nil
}

func (n *notificationService) StartScheduler() {
	err := scheduler.Start(n.run)
	if err != nil {
		logger.Errorf("can`t start scheduler:  %s\n", err.Error())
		return
	}

	n.chTasks = make(chan domain.Task, 100)
	for i := 0; i < countGoroutines; i++ {
		go n.worker()
	}

	now := time.Now().Truncate(time.Minute)
	n.fromTime = now.Unix()
	n.toTime = now.Add(time.Minute).Unix() - 1

	n.run()
}

func (n *notificationService) run() {
	tasks, err := n.repo.GetTasks(n.fromTime, n.toTime)
	if err != nil {
		logger.Info("can`t get task list")
		//если не смогли получить задачи из бд, следующий раз забираем задачи этой минуты
		n.toTime = time.Unix(n.toTime, 0).Add(time.Minute).Unix()
		return
	}

	n.fromTime = time.Unix(n.toTime+1, 0).Unix()
	n.toTime = time.Unix(n.toTime, 0).Add(time.Minute).Unix()

	for _, task := range tasks {
		n.chTasks <- task
	}
}

func (n *notificationService) worker() {
	for task := range n.chTasks {
		err := n.send.SendNotification(task.Email, task.Name)
		if err != nil {
			logger.Info("can`t send notification")
			continue
		}
		err = n.repo.Delete(task.ID)
		if err != nil {
			logger.Info("can`t delete task from db")
		}
	}
}
