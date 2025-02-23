package commands_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yvv4git/go-safe-upd/internal/infrastructure/commands"
)

func TestFetchVersionsList(t *testing.T) {
	const testModule = "github.com/jackc/pgx/v5"

	versions, err := commands.FetchVersionsList(testModule)
	require.NoError(t, err, "FetchVersionsList should not return an error")

	var actualVersions []string
	for version := range versions {
		actualVersions = append(actualVersions, version)
	}

	require.NotEmpty(t, actualVersions, "Versions list should not be empty")
}
