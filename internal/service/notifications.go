package service

import (
	"context"
	"github.com/mrcelviano/notificationservice/internal/domain"
	"github.com/mrcelviano/notificationservice/pkg/logger"
	"github.com/mrcelviano/notificationservice/pkg/scheduler"
	"github.com/mrcelviano/notificationservice/pkg/user"
	"time"
)

const (
	timeValue       = time.Minute * 1
	countGoroutines = 5
)

type notificationService struct {
	repositoryPG domain.NotificationRepositoryPG
	user         user.Service
	sender       domain.SenderService

	chTasks  chan domain.Task
	fromTime int64
	toTime   int64
}

func NewNotificationService(repositoryPG domain.NotificationRepositoryPG, sender domain.SenderService, user user.Service) domain.NotificationService {
	return &notificationService{
		repositoryPG: repositoryPG,
		sender:       sender,
		user:         user,
	}
}

func (n *notificationService) RegisterTask(ctx context.Context, userID int64) (bool, error) {
	err := n.repositoryPG.Create(ctx,
		domain.Task{
			RunTime: time.Now().Add(timeValue).Unix(),
			UserID:  userID,
		})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (n *notificationService) StartNotificationScheduler() {
	err := scheduler.Start(n.run)
	if err != nil {
		logger.Errorf("can`t start scheduler: %s\n", err.Error())
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
	tasks, err := n.repositoryPG.GetTasks(n.fromTime, n.toTime)
	if err != nil {
		logger.Errorf("can`t get task list: %s\n", err.Error())
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
		ctx := context.Background()
		userByID, err := n.user.GetUserByID(ctx, task.UserID)
		if err != nil {
			logger.Errorf("can`t get userByID: %s\n", err.Error())
			continue
		}
		err = n.sender.SendNotification(userByID.Email, userByID.Name)
		if err != nil {
			logger.Errorf("can`t sender notification: %s\n", err.Error())
			continue
		}
		err = n.repositoryPG.Delete(task.ID)
		if err != nil {
			logger.Errorf("can`t delete task: %s\n", err.Error())
			continue
		}
		isSetStatus, err := n.user.SetIsRegisteredStatus(ctx, task.UserID)
		if err != nil && !isSetStatus {
			logger.Errorf("can`t set status registered user: %s\n", err.Error())
		}
	}
}
