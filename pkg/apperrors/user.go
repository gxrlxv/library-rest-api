package apperrors

import (
	"errors"
)

var (
	ErrUserAlreadyExistEmail    = errors.New("user with given email already exist")
	ErrUserAlreadyExistName     = errors.New("user with given username already exist")
	ErrUserNotFound             = errors.New("user not found")
	ErrUserPasswordNotGenerated = errors.New("password generation error")
	ErrUserIncorrectPassword    = errors.New("incorrect password")
)
