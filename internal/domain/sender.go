package domain

type SenderService interface {
	SendNotification(string, string) error
}
