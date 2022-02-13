package user

import "errors"

var (
	ErrUserAddressNotFound  = errors.New("user service address not found")
	ErrUserNotFound         = errors.New("user is not found")
	ErrNotSetRegisterStatus = errors.New("can`t set status registered from user")
)
