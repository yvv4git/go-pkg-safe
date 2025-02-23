package commands

import (
	"fmt"
	"iter"
	"os/exec"
	"strings"
)

func FetchVersionsList(module string) (iter.Seq[string], error) {
	cmd := exec.Command("go", "list", "-m", "-versions", module)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("fetch versions: %w", err)
	}

	lines := strings.Fields(string(output))
	if len(lines) < 2 {
		return nil, fmt.Errorf("no versions found for module: %s", module)
	}

	versions := lines[1:] // skip the first line, which is the module path

	return func(yield func(string) bool) {
		for i := len(versions) - 1; i >= 0; i-- {
			if !yield(versions[i]) {
				return
			}
		}
	}, nil
}
