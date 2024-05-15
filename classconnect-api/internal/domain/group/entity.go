package group

import (
	"errors"
	"time"
)

var (
	ErrNotFound                  = errors.New("group not found")
	ErrWrongCode                 = errors.New("wrong group code")
	ErrAlreadyExistsSubWithGroup = errors.New("subscriber has group")
	ErrAlreadyExists             = errors.New("group already exists with this name")
)

type Group struct {
	ID               uint64
	Name             string
	Code             string
	IsExistsSchedule bool
	MembersCount     uint
	CreatedAt        time.Time
}
