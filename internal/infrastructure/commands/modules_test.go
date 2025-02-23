package commands_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yvv4git/go-pkg-safe/internal/infrastructure/commands"
)

func TestFetchModules(t *testing.T) {
	const currentModule = "github.com/yvv4git/go-pkg-safe"

	modules, err := commands.FetchModules()
	require.NoError(t, err, "FetchModules should not return an error")

	var actualModules []string
	for module := range modules {
		actualModules = append(actualModules, module)
	}

	require.Contains(t, actualModules, currentModule, "Current module %s should be in the list of modules", currentModule)
}
