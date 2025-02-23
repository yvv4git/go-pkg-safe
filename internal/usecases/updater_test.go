package usecases

import (
	"context"
	"iter"
	"testing"
	"time"

	"github.com/yvv4git/go-pkg-safe/internal/infrastructure/logger"
)

func TestUpdater(t *testing.T) {
	upd := NewUpdater(
		logger.SetupDefaultLogger(),
		ParamsNewUpdater{
			FetchModules: func() (iter.Seq[string], error) {
				stubModules := []string{
					"github.com/davecgh/go-spew",
					"github.com/stretchr/testify",
					"gopkg.in/yaml.v3",
				}

				return func(yield func(string) bool) {
					for _, modeName := range stubModules {
						if !yield(modeName) {
							return
						}
					}
				}, nil
			},
			FetchVersions: func(module string) (iter.Seq[string], error) {
				stubVersions := []string{
					"v1.3.0",
					"v1.2.0",
					"v1.1.0",
				}

				return func(yield func(string) bool) {
					for _, version := range stubVersions {
						if !yield(version) {
							return
						}
					}
				}, nil
			},
			IsSafityModuleVersion: func(module, version string, timeThreshold time.Duration) (bool, error) {
				if module == "github.com/davecgh/go-spew" && version == "v1.1.0" {
					return true, nil
				}

				return false, nil
			},
		},
	)
	err := upd.Update(context.Background(), time.Hour*24*7)
	t.Logf("err: %v", err)
}
