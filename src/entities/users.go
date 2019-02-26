package entities

import (
	"errors"
)

var (
	ErrUserNotfound = errors.New("user cannot be found")
)

type Users struct {
	ID    int    `json:"user_id"`
	Email string `json:"email"`
}
