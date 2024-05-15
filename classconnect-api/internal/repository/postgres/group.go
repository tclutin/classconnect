package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/tclutin/classconnect-api/internal/domain/group"
	"github.com/tclutin/classconnect-api/pkg/client/postgresql"
	"log/slog"
)

const (
	layerGroupRepository = "repository.user."
)

type GroupRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewGroupRepository(client postgresql.Client, logger *slog.Logger) *GroupRepository {
	return &GroupRepository{
		db:     client,
		logger: logger,
	}
}

func (g *GroupRepository) GetGroupById(ctx context.Context, id string) (group.Group, error) {
	sql := `SELECT * FROM public.groups WHERE id = $1`

	var getGroup group.Group

	row := g.db.QueryRow(ctx, sql, id)

	err := row.Scan(
		&getGroup.ID,
		&getGroup.Name,
		&getGroup.Code,
		&getGroup.IsExistsSchedule,
		&getGroup.MembersCount,
		&getGroup.CreatedAt)

	if err != nil {
		return group.Group{}, err
	}

	return getGroup, nil
}

func (g *GroupRepository) CreateGroup(ctx context.Context, group group.Group) error {
	sql := `INSERT INTO public.groups (name, code, created_at) VALUES ($1, $2, $3)`

	_, err := g.db.Exec(ctx, sql, group.Name, group.Code, group.CreatedAt)

	return err
}

func (g *GroupRepository) GetGroupByName(ctx context.Context, name string) (group.Group, error) {
	sql := `SELECT * FROM public.groups WHERE name = $1`

	var getGroup group.Group

	row := g.db.QueryRow(ctx, sql, name)

	err := row.Scan(
		&getGroup.ID,
		&getGroup.Name,
		&getGroup.Code,
		&getGroup.IsExistsSchedule,
		&getGroup.MembersCount,
		&getGroup.CreatedAt)

	if err != nil {
		return group.Group{}, err
	}

	return getGroup, nil
}

func (g *GroupRepository) GetAllGroups(ctx context.Context) ([]group.Group, error) {
	sql := `SELECT id, name, is_exists_schedule, members_count, created_at FROM public.groups`

	rows, err := g.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	groups := make([]group.Group, 0)

	for rows.Next() {
		var getGroup group.Group

		err = rows.Scan(
			&getGroup.ID,
			&getGroup.Name,
			&getGroup.IsExistsSchedule,
			&getGroup.MembersCount,
			&getGroup.CreatedAt)

		if err != nil {
			return nil, err
		}

		groups = append(groups, getGroup)
	}

	return groups, nil
}

func (g *GroupRepository) UpdateGroup(ctx context.Context, group group.Group) error {
	sql := `UPDATE public.groups SET name = $1,
                         code = $2,
                         is_exists_schedule = $3,
                         members_count = $4 WHERE id = $5`

	_, err := g.db.Exec(ctx, sql, group.Name, group.Code, group.IsExistsSchedule, group.MembersCount, group.ID)

	return err
}

// DeleteGroup with transactions
func (g *GroupRepository) DeleteGroup(ctx context.Context, groupID uint64) error {
	tx, err := g.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err = g.untieSubscribers(ctx, tx, groupID); err != nil {
		return err
	}

	if err = g.untieUser(ctx, tx, groupID); err != nil {
		return err
	}

	sql := `DELETE FROM public.groups WHERE id = $1`

	_, err = tx.Exec(ctx, sql, groupID)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return err
}

func (g *GroupRepository) untieSubscribers(ctx context.Context, tx pgx.Tx, groupID uint64) error {
	sql := `UPDATE public.subscribers SET group_id = NULL WHERE group_id = $1`

	_, err := tx.Exec(ctx, sql, groupID)

	return err
}

func (g *GroupRepository) untieUser(ctx context.Context, tx pgx.Tx, groupID uint64) error {
	sql := `UPDATE public.users	 SET group_id = NULL WHERE group_id = $1`

	_, err := tx.Exec(ctx, sql, groupID)

	return err
}
