package entity

import "errors"

var (
	ErrPersonNotFound = errors.New("person not found")
	ErrGroupNotFound  = errors.New("group not found")
)
