package apperrors

import (
	"errors"
)

var (
	ErrUserAlreadyExistEmail    = errors.New("user with given email already exist")
	ErrUserAlreadyExistName     = errors.New("user with given username already exist")
	ErrUserNotFound             = errors.New("user not found")
	ErrUserNotArchived          = errors.New("user cannot be archived")
	ErrUserContextNotFound      = errors.New("user not found in context")
	ErrUserPasswordNotGenerated = errors.New("password generation error")
	ErrUserIncorrectEmail       = errors.New("incorrect email")
	ErrUserIncorrectName        = errors.New("incorrect username")
	ErrUserIncorrectPassword    = errors.New("incorrect password")
)
