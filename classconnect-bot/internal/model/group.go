package model

import "time"

type JoinToGroupRequest struct {
	ID   uint64 `json:"sub_id"`
	Code string `json:"code"`
}

type LeaveFromGroupRequest struct {
	ID uint64 `json:"sub_id"`
}

type GroupResponse struct {
	ID               uint64    `json:"id"`
	Name             string    `json:"name"`
	IsExistsSchedule bool      `json:"is_exists_schedule"`
	MembersCount     uint      `json:"members_count"`
	CreatedAt        time.Time `json:"created_at"`
}
