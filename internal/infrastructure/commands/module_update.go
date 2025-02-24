package commands

import (
	"fmt"

	"os/exec"
)

func ModuleUpdate(module, version string) error {
	cmd := exec.Command("go", "get", "-v", "-u", module+"@"+version)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("update module %s to version %s: %w", module, version, err)
	}

	return nil
}
