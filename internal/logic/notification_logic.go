package logic

import "github.com/mrcelviano/notificationservice/app"

type notificationLogic struct {
	repo app.NotificationRepository
}

func NewNotificationLogic(repo app.NotificationRepository) app.NotificationLogic {
	return &notificationLogic{repo: repo}
}
