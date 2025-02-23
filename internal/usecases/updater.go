package usecases

import (
	"context"
	"fmt"
	"iter"
	"log/slog"
	"time"
)

type (
	fnFetchModules          func() (iter.Seq[string], error)
	fnFetchVersions         func(module string) (iter.Seq[string], error)
	fnIsSafityModuleVersion func(module, version string, timeThreshold time.Duration) (bool, error)
	fnModuleUpdate          func(module, version string) error
)

type Updater struct {
	log                   *slog.Logger
	fetchModules          fnFetchModules
	fetchVersions         fnFetchVersions
	isSafityModuleVersion fnIsSafityModuleVersion
	moduleUpdate          fnModuleUpdate
}

type ParamsNewUpdater struct {
	FetchModules          fnFetchModules
	FetchVersions         fnFetchVersions
	IsSafityModuleVersion fnIsSafityModuleVersion
	ModuleUpdate          fnModuleUpdate
}

func NewUpdater(log *slog.Logger, params ParamsNewUpdater) *Updater {
	return &Updater{
		log:                   log,
		fetchModules:          params.FetchModules,
		fetchVersions:         params.FetchVersions,
		isSafityModuleVersion: params.IsSafityModuleVersion,
		moduleUpdate:          params.ModuleUpdate,
	}
}

func (u *Updater) Update(ctx context.Context, vertionThreshold time.Duration) error {
	u.log.Info("Updating...")

	modules, err := u.fetchModules()
	if err != nil {
		return fmt.Errorf("fetch modules: %w", err)
	}

	for module := range modules {
		if ctx.Err() != nil {
			return fmt.Errorf("context canceled: %w", ctx.Err())
		}

		if err := u.processModuleVersions(ctx, module, vertionThreshold); err != nil {
			return fmt.Errorf("process module[%s] versions: %w", module, err)
		}
	}

	return nil
}

func (u *Updater) processModuleVersions(ctx context.Context, module string, versionThreshold time.Duration) error {
	versions, err := u.fetchVersions(module)
	if err != nil {
		return err
	}
	u.log.Info("Versions fetched", "module", module)

	for version := range versions {
		if ctx.Err() != nil {
			return fmt.Errorf("context canceled: %w", ctx.Err())
		}

		isSafity, err := u.isSafityModuleVersion(module, version, versionThreshold)
		if err != nil {
			u.log.Error("Failed to check if version is older than threshold", "module", module, "version", version, "err", err)
			continue
		}

		if isSafity {
			u.log.Info("Found a safity version", "module", module, "version", version)
			return u.moduleUpdate(module, version)
		}

		u.log.Info("Version is not safity", "module", module, "version", version)
	}

	return nil
}
