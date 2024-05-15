package auth

import "time"

type LoginDTO struct {
	Username string
	Password string
}

type SignupDTO struct {
	Username string
	Email    string
	Password string
}

type UserDetailDTO struct {
	ID       uint64
	Username string
	Email    string
	IsBanned bool
	Group    *UserGroupDTO
}

type UserGroupDTO struct {
	ID               uint64
	Name             string
	Code             string
	IsExistsSchedule bool
	MembersCount     uint
	CreatedAt        time.Time
}
