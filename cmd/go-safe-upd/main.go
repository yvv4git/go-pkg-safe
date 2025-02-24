package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/yvv4git/go-safe-upd/internal/infrastructure/commands"
	"github.com/yvv4git/go-safe-upd/internal/infrastructure/logger"
	"github.com/yvv4git/go-safe-upd/internal/usecases"
)

func main() {
	log := logger.SetupLoggerWithLevel(slog.LevelInfo)

	versionThresholdDays := kingpin.Flag("version-threshold", "Threshold in days for module versions").
		Short('t').
		Default("14").
		Int()

	kingpin.Parse()

	versionThreshold := time.Duration(*versionThresholdDays) * 24 * time.Hour

	ucUpdater := usecases.NewUpdater(log, usecases.ParamsNewUpdater{
		FetchModules:          commands.FetchModules,
		FetchVersions:         commands.FetchVersionsList,
		IsSafityModuleVersion: commands.IsSafityModuleVersion,
		ModuleUpdate:          commands.ModuleUpdate,
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := ucUpdater.Update(ctx, versionThreshold); err != nil {
		log.Error(err.Error())
	}

	log.Info("All modules updated")
}
