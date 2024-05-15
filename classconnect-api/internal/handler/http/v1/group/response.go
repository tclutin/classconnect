package group

import (
	"github.com/tclutin/classconnect-api/internal/domain/group"
	"time"
)

type GroupResponse struct {
	ID               uint64    `json:"id"`
	Name             string    `json:"name"`
	IsExistsSchedule bool      `json:"is_exists_schedule"`
	MembersCount     uint      `json:"members_count"`
	CreatedAt        time.Time `json:"created_at"`
}

func ConvertGroupToResponse(group group.Group) GroupResponse {
	return GroupResponse{
		ID:               group.ID,
		Name:             group.Name,
		IsExistsSchedule: group.IsExistsSchedule,
		MembersCount:     group.MembersCount,
		CreatedAt:        group.CreatedAt,
	}
}

func ConvertGroupsToResponse(groups []group.Group) []GroupResponse {
	response := make([]GroupResponse, 0, len(groups))

	for _, value := range groups {
		groupResponse := ConvertGroupToResponse(value)
		response = append(response, groupResponse)
	}

	return response
}
