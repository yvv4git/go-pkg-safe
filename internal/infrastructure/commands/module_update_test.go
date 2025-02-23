package commands_test

import (
	"testing"

	"github.com/yvv4git/go-pkg-safe/internal/infrastructure/commands"
)

func TestModuleUpdate(t *testing.T) {
	const (
		moduleName    = "github.com/stretchr/testify"
		moduleVersion = "v1.10.0"
	)

	err := commands.ModuleUpdate(moduleName, moduleVersion)
	t.Logf("err: %v", err)
}
