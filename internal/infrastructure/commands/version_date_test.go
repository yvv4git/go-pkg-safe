package commands_test

import (
	"testing"
	"time"

	"github.com/test-go/testify/assert"
	"github.com/yvv4git/go-pkg-safe/internal/infrastructure/commands"
)

func TestFetchVersionDate(t *testing.T) {
	const (
		moduleName    = "github.com/jackc/pgx/v5"
		moduleVersion = "v5.7.2"
		timeThreshold = 14 * 24 * time.Hour // 2 weeks
	)

	isOlder, err := commands.IsSafityModuleVersion(moduleName, moduleVersion, timeThreshold)
	assert.NoError(t, err)
	assert.True(t, isOlder)
}
