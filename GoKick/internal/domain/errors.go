package domain

import "errors"

var (
	ErrNotFound  = errors.New("campaign not found")
	ErrDuplicate = errors.New("campaign with given ID already exists")
)
