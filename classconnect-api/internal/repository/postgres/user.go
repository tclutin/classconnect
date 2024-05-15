package postgres

import (
	"context"
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"github.com/tclutin/classconnect-api/pkg/client/postgresql"
	"log/slog"
)

const (
	layerUserRepository = "repository.user."
)

type UserRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewUserRepository(client postgresql.Client, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		db:     client,
		logger: logger,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, user auth.User) error {
	sql := `INSERT INTO public.users (username, email, hashed_password, created_at) VALUES ($1, $2, $3, $4)`

	_, err := u.db.Exec(ctx, sql, user.Username, user.Email, user.PasswordHash, user.CreatedAt)

	return err
}

// TODO: one day need to fix this cringe
func (u *UserRepository) GetUserByUsernameWithDetail(ctx context.Context, username string) (auth.UserDetailDTO, error) {
	sql := `SELECT u.id, u.username, u.email, u.is_banned, COALESCE(g.id, 0), COALESCE(g.name, ''),COALESCE(g.code, ''), COALESCE(g.is_exists_schedule, false), COALESCE(g.members_count, 0), COALESCE(g.created_at, NOW())
			FROM public.users u
			LEFT JOIN public.groups g ON u.group_id = g.id
			WHERE u.username = $1`

	var userDetail auth.UserDetailDTO
	userDetail.Group = &auth.UserGroupDTO{}

	row := u.db.QueryRow(ctx, sql, username)
	err := row.Scan(
		&userDetail.ID,
		&userDetail.Username,
		&userDetail.Email,
		&userDetail.IsBanned,
		&userDetail.Group.ID,
		&userDetail.Group.Name,
		&userDetail.Group.Code,
		&userDetail.Group.IsExistsSchedule,
		&userDetail.Group.MembersCount,
		&userDetail.Group.CreatedAt,
	)
	if err != nil {
		return auth.UserDetailDTO{}, err
	}

	if userDetail.Group.ID == 0 && userDetail.Group.Name == "" {
		userDetail.Group = nil
	}

	return userDetail, nil
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (auth.User, error) {
	sql := `SELECT * FROM public.users WHERE username = $1`

	u.logger.Info(layerUserRepository+"GetUserByUsername", slog.String("sql", sql))

	var user auth.User

	row := u.db.QueryRow(ctx, sql, username)

	err := row.Scan(
		&user.ID,
		&user.GroupID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.IsBanned,
		&user.CreatedAt)

	if err != nil {
		return auth.User{}, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user auth.User) error {
	sql := `UPDATE public.users SET group_id = $1,
                        username = $2,
                        email = $3,
                        hashed_password = $4,
                        is_banned = $5 WHERE id = $6`

	_, err := u.db.Exec(ctx, sql,
		user.GroupID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.IsBanned,
		user.ID)

	return err
}
