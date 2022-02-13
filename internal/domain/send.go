package domain

type SendService interface {
	SendNotification(string, string) error
}
