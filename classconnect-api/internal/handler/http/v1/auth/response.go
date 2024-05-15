package auth

import (
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type UserDetailResponse struct {
	ID       uint64             `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	IsBanned bool               `json:"is_banned"`
	Group    *UserGroupResponse `json:"group"`
}

type UserGroupResponse struct {
	ID               uint64    `json:"id"`
	Name             string    `json:"name"`
	Code             string    `json:"code"`
	IsExistsSchedule bool      `json:"is_exists_schedule"`
	MembersCount     uint      `json:"members_count"`
	CreatedAt        time.Time `json:"created_at"`
}

func ConvertUserDetailDTOToResponse(userDetail auth.UserDetailDTO) UserDetailResponse {
	if userDetail.Group == nil {
		return UserDetailResponse{
			ID:       userDetail.ID,
			Username: userDetail.Username,
			Email:    userDetail.Email,
			IsBanned: userDetail.IsBanned,
			Group:    nil,
		}
	}

	group := &UserGroupResponse{
		ID:               userDetail.Group.ID,
		Name:             userDetail.Group.Name,
		Code:             userDetail.Group.Code,
		IsExistsSchedule: userDetail.Group.IsExistsSchedule,
		MembersCount:     userDetail.Group.MembersCount,
		CreatedAt:        userDetail.Group.CreatedAt,
	}

	return UserDetailResponse{
		ID:       userDetail.ID,
		Username: userDetail.Username,
		Email:    userDetail.Email,
		IsBanned: userDetail.IsBanned,
		Group:    group,
	}
}
