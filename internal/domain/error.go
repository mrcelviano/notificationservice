package domain

import "errors"

var (
	ErrRegisterNotification = errors.New("failed to register notification")
	ErrCantExecSQLRequest   = errors.New("cant exec sql request")
)
