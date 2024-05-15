package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/tclutin/classconnect-api/internal/domain/schedule"
	"github.com/tclutin/classconnect-api/pkg/client/postgresql"
	"log/slog"
	"time"
)

const (
	layerScheduleRepository = "repository.schedule."
)

type ScheduleRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewScheduleRepository(client postgresql.Client, logger *slog.Logger) *ScheduleRepository {
	return &ScheduleRepository{
		db:     client,
		logger: logger,
	}
}

func (s *ScheduleRepository) GetScheduleForDay(ctx context.Context, groupID uint64, dayNumber int, isEven bool) ([]schedule.SubjectDTO, error) {
	sql := `SELECT s.teacher, s.name, s.cabinet, s.description, s.time_start, s.time_end FROM public.subjects AS s
    		INNER JOIN public.days as d ON d.id = s.day_id
    		INNER JOIN public.weeks as w ON w.id = d.week_id
    		INNER JOIN public.groups as g ON g.id = w.group_id
    		WHERE d.day_of_week = $1 AND w.is_even = $2 AND g.id = $3
    		ORDER BY s.time_start`

	rows, err := s.db.Query(ctx, sql, dayNumber, isEven, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []schedule.SubjectDTO

	for rows.Next() {
		var subject schedule.SubjectDTO
		var startTime, endTime time.Time
		err = rows.Scan(&subject.Teacher,
			&subject.Name,
			&subject.Cabinet,
			&subject.Description,
			&startTime,
			&endTime)

		if err != nil {
			return nil, err
		}

		subject.StartTime = startTime.Format("15:04")
		subject.EndTime = endTime.Format("15:04")

		subjects = append(subjects, subject)
	}

	return subjects, nil
}

// CreateSchedule with transactions
func (s *ScheduleRepository) CreateSchedule(ctx context.Context, schedule schedule.UploadScheduleDTO, groupID uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, week := range schedule.Weeks {
		weekID, err := s.createWeek(ctx, tx, week, groupID)
		if err != nil {
			return err
		}

		for _, day := range week.Days {
			dayID, err := s.createDay(ctx, tx, weekID, day)
			if err != nil {
				return err
			}

			for _, subject := range day.Subjects {
				err := s.createSubject(ctx, tx, dayID, subject)
				if err != nil {
					return err
				}
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) createWeek(ctx context.Context, tx pgx.Tx, week schedule.WeekDTO, groupID uint64) (uint64, error) {
	sql := `INSERT INTO public.weeks (group_id, is_even) VALUES ($1, $2) RETURNING id`

	var weekID uint64

	err := tx.QueryRow(ctx, sql, groupID, week.IsEven).Scan(&weekID)
	if err != nil {
		return 0, err
	}

	return weekID, nil
}

func (s *ScheduleRepository) createDay(ctx context.Context, tx pgx.Tx, weekID uint64, day schedule.DayDTO) (uint64, error) {
	sql := `INSERT INTO public.days (week_id, day_of_week) VALUES ($1, $2) RETURNING id`

	var dayID uint64

	err := tx.QueryRow(ctx, sql, weekID, day.DayNumber).Scan(&dayID)
	if err != nil {
		return 0, err
	}

	return dayID, nil
}

func (s *ScheduleRepository) createSubject(ctx context.Context, tx pgx.Tx, dayID uint64, day schedule.SubjectDTO) error {
	sql := `INSERT INTO public.subjects (day_id,
                             teacher,
                             name,
                             cabinet,
                             description,
                             time_start,
                             time_end) VALUES ($1,$2, $3, $4, $5, $6, $7)`

	_, err := tx.Exec(ctx, sql,
		dayID,
		day.Teacher,
		day.Name,
		day.Cabinet,
		day.Description,
		day.StartTime,
		day.EndTime)

	if err != nil {
		return err
	}

	return nil
}

// DeleteSchedule with transcations
func (s *ScheduleRepository) DeleteSchedule(ctx context.Context, groupID uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err = s.deleteSubjects(ctx, tx, groupID); err != nil {
		return err
	}

	if err = s.deleteDays(ctx, tx, groupID); err != nil {
		return err
	}

	if err = s.deleteWeeks(ctx, tx, groupID); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) deleteSubjects(ctx context.Context, tx pgx.Tx, groupID uint64) error {
	sql := `DELETE FROM public.subjects WHERE day_id IN (SELECT id FROM public.days WHERE week_id IN (SELECT id FROM public.weeks WHERE group_id = $1))`

	_, err := tx.Exec(ctx, sql, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) deleteDays(ctx context.Context, tx pgx.Tx, groupID uint64) error {
	sql := `DELETE FROM public.days WHERE week_id IN (SELECT id FROM public.weeks WHERE group_id = $1)`

	_, err := tx.Exec(ctx, sql, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) deleteWeeks(ctx context.Context, tx pgx.Tx, groupID uint64) error {
	sql := `DELETE FROM public.weeks WHERE group_id = $1`

	_, err := tx.Exec(ctx, sql, groupID)
	if err != nil {
		return err
	}

	return nil
}
