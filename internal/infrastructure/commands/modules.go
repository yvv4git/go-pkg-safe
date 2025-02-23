package commands

import (
	"fmt"
	"iter"
	"os/exec"
	"strings"
)

func FetchModules() (iter.Seq[string], error) {
	cmd := exec.Command("go", "list", "-m", "-u", "all")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("fetch modules: %w", err)
	}

	lines := strings.Split(string(output), "\n")

	return func(yield func(string) bool) {
		for _, line := range lines {
			if line == "" || !strings.Contains(line, "[") {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) < 3 {
				continue
			}

			moduleName := parts[0] // get module name

			if !yield(moduleName) {
				return
			}
		}
	}, nil
}
