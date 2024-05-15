package repository

import (
	"classconnect-notification-service/internal/config"
	"classconnect-notification-service/internal/entity"
	"classconnect-notification-service/pkg/client/postgresql"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type ScheduleRepository struct {
	client postgresql.Client
	logger *slog.Logger
	cfg    *config.Config
}

func NewScheduleRepository(client postgresql.Client, cfg *config.Config, logger *slog.Logger) *ScheduleRepository {
	return &ScheduleRepository{client: client, cfg: cfg, logger: logger}
}

func (s *ScheduleRepository) GetSubjectsWithDetail(ctx context.Context, dayOfweek int, isEven bool, currentTime time.Time) ([]entity.Subject, error) {
	sql := `SELECT s.teacher, s.name, s.cabinet, s.time_start, s.time_end, sub.device_id, sub.telegram_chat_id
			FROM public.subjects AS s
			INNER JOIN public.days AS d ON s.day_id = d.id
			INNER JOIN public.weeks AS w ON d.week_id = w.id
			INNER JOIN public.subscribers AS sub ON sub.group_id = w.group_id
			WHERE d.day_of_week = $1 AND w.is_even = $2 AND s.time_start BETWEEN $3 AND $4`

	nextTime := currentTime.Add(time.Duration(s.cfg.InDelay) * time.Minute)

	currentTimeFormatted := currentTime.Format("15:04")
	nextTimeFormatted := nextTime.Format("15:04")

	rangeOftime := fmt.Sprintf("%v - %v", currentTimeFormatted, nextTimeFormatted)

	s.logger.Debug("getting range of time", slog.String("range", rangeOftime))

	rows, err := s.client.Query(ctx, sql, dayOfweek, isEven, currentTime, nextTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []entity.Subject

	for rows.Next() {
		var subject entity.Subject
		var startTime, endTime time.Time
		err = rows.Scan(
			&subject.Teacher,
			&subject.Name,
			&subject.Cabinet,
			&startTime,
			&endTime,
			&subject.DeviceID,
			&subject.ChatID)

		if err != nil {
			return nil, err
		}

		subject.StartTime = startTime.Format("15:04")
		subject.EndTime = endTime.Format("15:04")

		subjects = append(subjects, subject)
	}

	return subjects, nil
}
