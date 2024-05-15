package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/tclutin/classconnect-api/internal/config"
	"github.com/tclutin/classconnect-api/internal/domain"
	httpLayer "github.com/tclutin/classconnect-api/internal/handler/http"
	"github.com/tclutin/classconnect-api/internal/repository/postgres"
	"github.com/tclutin/classconnect-api/pkg/client/postgresql"
	"github.com/tclutin/classconnect-api/pkg/logging"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New() *App {
	//Init the config
	cfg := config.MustLoad()

	//Init the slog
	logger := logging.InitSlog(cfg.Environment)

	connectStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DbName)

	//Init postgres client
	client := postgresql.NewClient(context.Background(), connectStr)

	//Init repositories
	repositories := postgres.NewRepositories(client, logger)

	//Init manager of serviсes
	services := domain.NewServices(cfg, repositories)

	//Init the router
	router := httpLayer.NewRouter(services, cfg, logger)

	return &App{
		logger: logger,
		httpServer: &http.Server{
			Addr:    net.JoinHostPort(cfg.HTTPServer.Address, cfg.HTTPServer.Port),
			Handler: router,
		},
	}
}

func (a *App) Run(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("http server closed unexpectedly", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	<-quit

	a.Stop(ctx)
}

func (a *App) Stop(ctx context.Context) {
	a.logger.Info("shutting down")
	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.logger.Error("an error occurred during server shutdown", slog.Any("error", err))
		os.Exit(1)
	}
}
