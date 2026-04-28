package domain

import "errors"

var (
	ErrTaskNotFound     = errors.New("task not found")
	ErrEmptyDescription = errors.New("description cannot be empty")
	ErrInvalidStatus    = errors.New("invalid status")
)
