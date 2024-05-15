package auth

import (
	"errors"
	"time"
)

var (
	ErrAlreadyExist     = errors.New("user already exist")
	ErrNotFound         = errors.New("user not found")
	ErrPasswordNotMatch = errors.New("password not match")
)

type User struct {
	ID           uint64
	GroupID      *uint64
	Username     string
	Email        string
	PasswordHash string
	IsBanned     bool
	CreatedAt    time.Time
}
