package user

import "context"

type User struct {
	ID    int64
	Email string
	Name  string
}

type Service interface {
	GetUserByID(context.Context, int64) (User, error)
	SetIsRegisteredStatus(context.Context, int64) (bool, error)
}
