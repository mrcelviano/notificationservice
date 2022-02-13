package service

import (
	"context"
	"github.com/mrcelviano/notificationservice/internal/domain"
	"github.com/pkg/errors"
	cron "github.com/robfig/cron/v3"
	"log"
	"time"
)

const (
	timeValue       = time.Minute * 30
	countGoroutines = 5
)

type notificationService struct {
	repo     domain.NotificationRepositoryPG
	send     domain.SendService
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
	runTime := time.Now()
	runTime = runTime.Add(timeValue)

	task.RunTime = runTime.Unix()
	task, err = n.repo.Create(ctx, task)
	if err != nil {
		return id, errors.Wrap(err, "can`t create task")
	}
	return task.ID, nil
}

func (n *notificationService) Start() {
	cr := cron.New(cron.WithSeconds())
	_, err := cr.AddFunc("0 * * * * *", n.run)
	if err != nil {
		panic(err)
	}

	cr.Start()
	n.chTasks = make(chan domain.Task, 100)
	for i := 0; i < countGoroutines; i++ {
		go n.worker()
	}

	now := time.Now().Truncate(time.Minute)
	n.fromTime = now.Unix()
	n.toTime = now.Add(time.Minute).Unix() - 1
	log.Println("Started")
	n.run()
}

func (n *notificationService) run() {
	tasks, err := n.repo.GetTasks(n.fromTime, n.toTime)
	if err != nil {
		log.Println("can`t get task list")
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
		n.send.SendNotification(task.Email, task.Name)
		if err := n.repo.Delete(task.ID); err != nil {
			log.Println("can`t delete task from db")
		}
	}
}
