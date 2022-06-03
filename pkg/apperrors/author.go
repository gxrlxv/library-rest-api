package apperrors

import "errors"

var (
	ErrAuthorAlreadyExistName = errors.New("author with given name already exist")
	ErrAuthorNotFound         = errors.New("author not found")
)
