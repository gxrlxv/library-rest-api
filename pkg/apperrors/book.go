package apperrors

import "errors"

var (
	ErrBookAlreadyExistName = errors.New("book with given name already exist")
	ErrBookNotFound         = errors.New("book not found")
)
