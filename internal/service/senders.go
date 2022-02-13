package service

import (
	"fmt"
	"github.com/mrcelviano/notificationservice/internal/domain"
)

type senderService struct{}

func NewSenderService() domain.SendService {
	return &senderService{}
}

func (s *senderService) SendNotification(email string, name string) {
	fmt.Println(fmt.Sprintf("Здравствуйте %v! Спасибо за регестрацию в нашей платформе. Ваш email: %v", name, email))
}
