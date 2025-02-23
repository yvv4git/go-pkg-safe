package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yvv4git/go-pkg-safe/internal/infrastructure/commands"
	"github.com/yvv4git/go-pkg-safe/internal/infrastructure/logger"
	"github.com/yvv4git/go-pkg-safe/internal/usecases"
)

const (
	defaultVersionThreshould = 14 * 24 * time.Hour
)

func main() {
	log := logger.SetupLoggerWithLevel(slog.LevelInfo)

	ucUpdater := usecases.NewUpdater(log, usecases.ParamsNewUpdater{
		FetchModules:          commands.FetchModules,
		FetchVersions:         commands.FetchVersionsList,
		IsSafityModuleVersion: commands.IsSafityModuleVersion,
		ModuleUpdate:          commands.ModuleUpdate,
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := ucUpdater.Update(ctx, defaultVersionThreshould); err != nil {
		log.Error(err.Error())
	}

	log.Info("All modules updated successfully")
}
