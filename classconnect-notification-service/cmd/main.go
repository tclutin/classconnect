package main

import (
	"classconnect-notification-service/internal/config"
	"classconnect-notification-service/internal/repository"
	"classconnect-notification-service/internal/service"
	"classconnect-notification-service/pkg/client/inmemory"
	"classconnect-notification-service/pkg/client/postgresql"
	"classconnect-notification-service/pkg/client/telegram"
	"classconnect-notification-service/pkg/logging"
	"classconnect-notification-service/pkg/utils"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {
	//Initializing the config
	cfg := config.MustLoad()

	//Initializing the custom slog
	opts := logging.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := logging.NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)

	connectStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName)

	//Initializing the postgresql client
	client := postgresql.NewClient(context.Background(), connectStr)

	//Initializing the telegram client
	tgClient := telegram.NewClient(cfg.Telegram.Token)

	//Initializing the notification service
	notificationService := service.NewNotificationService(tgClient, logger)

	//Initializing the schedule repository
	scheduleRepository := repository.NewScheduleRepository(client, cfg, logger)

	//Initializing the inmemory storage
	memoryStorage := inmemory.NewInMemoryStorage(50, 8)

	ticker := time.NewTicker(time.Duration(cfg.ExDelay) * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			currentTime := time.Now().UTC()
			subjects, err := scheduleRepository.GetSubjectsWithDetail(context.TODO(), utils.GetDayOfWeek(), utils.IsEvenWeek(), currentTime)
			if err != nil {
				slog.Error("an error occurred while receiving the data", slog.Any("error", err))
				continue
			}
			for _, subject := range subjects {
				if !memoryStorage.Get(*subject.ChatID) {
					remainedTime, err := utils.GetDifferentTimeRFFC822(subject.StartTime)
					if err != nil {
						slog.Error("an error occurred while parsing the date", slog.Any("error", err))
					}

					message := fmt.Sprintf("✅ Notification ✅\n\n📖 Subject: %s\n📍 Location: %s\n👨‍🏫 Teacher: %s\n⏰ Time: %s - %s\n⏳ Time Until: %d minutes", subject.Name, subject.Cabinet, subject.Teacher, subject.StartTime, subject.EndTime, remainedTime)
					notificationService.Send(*subject.ChatID, message)
					memoryStorage.Set(*subject.ChatID, true)
				}
			}
		}
	}

}
