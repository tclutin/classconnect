package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tclutin/classconnect-api/internal/config"
	"github.com/tclutin/classconnect-api/pkg/hash"
	"time"
)

const (
	layerAuthService = "service.auth."
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Username string
	Email    string
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByUsernameWithDetail(ctx context.Context, username string) (UserDetailDTO, error)
}

type Service struct {
	config     *config.Config
	repository Repository
}

func (s *Service) GetUserByUsernameWithDetail(ctx context.Context, username string) (UserDetailDTO, error) {
	return s.repository.GetUserByUsernameWithDetail(ctx, username)
}

func NewService(config *config.Config, repository Repository) *Service {
	return &Service{
		config:     config,
		repository: repository,
	}
}

func (s *Service) LogIn(ctx context.Context, dto LoginDTO) (string, error) {
	user, err := s.repository.GetUserByUsername(ctx, dto.Username)
	if err != nil {
		return "", ErrNotFound
	}

	if user.PasswordHash != hash.GenerateSha1Hash(dto.Password) {
		return "", ErrPasswordNotMatch
	}

	token, err := s.GenerateToken(user.Username, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) SignUp(ctx context.Context, dto SignupDTO) (string, error) {
	if _, err := s.repository.GetUserByUsername(ctx, dto.Username); err == nil {
		return "", ErrAlreadyExist
	}

	user := User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: hash.GenerateSha1Hash(dto.Password),
		CreatedAt:    time.Now(),
	}

	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(user.Username, dto.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (User, error) {
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return User{}, ErrNotFound
	}

	return user, nil
}

func (s *Service) GenerateToken(username string, email string) (string, error) {
	duration, err := time.ParseDuration(s.config.JWT.Expire)
	if err != nil {
		return "", errors.New("error of parsing duration")
	}

	payload := jwt.MapClaims{
		"exp":      time.Now().Add(duration).Unix(),
		"username": username,
		"email":    email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) ParseToken(accessToken string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}
