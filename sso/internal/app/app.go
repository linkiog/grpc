package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/linkiog/sso/internal/app/grpc"
	"github.com/linkiog/sso/internal/services/auth"
	storage "github.com/linkiog/sso/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTl time.Duration,

) *App {
	storage, err := storage.New(storagePath)
	if err != nil {
		panic(err)

	}
	authService := auth.New(log, storage, storage, storage, tokenTTl)
	grpcApp := grpcapp.New(log, authService, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}

}
