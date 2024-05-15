package auth

import "github.com/tclutin/classconnect-api/internal/domain/auth"

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

type SignupRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Email    string `json:"email" binding:"required,email,max=40"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

func (l LoginRequest) ToDTO() auth.LoginDTO {
	return auth.LoginDTO{
		Username: l.Username,
		Password: l.Password,
	}
}

func (l SignupRequest) ToDTO() auth.SignupDTO {
	return auth.SignupDTO{
		Username: l.Username,
		Email:    l.Email,
		Password: l.Password,
	}
}
