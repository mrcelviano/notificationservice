package repository

import "github.com/mrcelviano/notificationservice/app"

type notificationRepository struct{}

func NewNotificationRepository() app.NotificationRepository {
	return &notificationRepository{}
}
