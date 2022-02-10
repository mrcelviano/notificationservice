package logic

import (
	"fmt"
	"github.com/mrcelviano/notificationservice/app"
)

type sendLogic struct{}

func NewSendNotificationLogic() app.SendLogic {
	return &sendLogic{}
}

func (s *sendLogic) SendNotification(email string, name string) {
	fmt.Println(fmt.Sprintf("Здравствуйте %v! Спасибо за регестрацию в нашей платформе. Ваш email: %v", name, email))
}
