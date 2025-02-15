package errors

import "errors"

var (
	ErrOnServer          = errors.New("error on Server")
	ErrUserAlreadyExists = errors.New("User already exists")
)
