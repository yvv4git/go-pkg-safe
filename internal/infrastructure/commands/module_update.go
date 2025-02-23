package commands

import (
	"fmt"

	"os/exec"
)

func ModuleUpdate(module, version string) error {
	cmd := exec.Command("go", "get", "-v", "-u", module+"@"+version)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("update module %s to version %s: %w", module, version, err)
	}

	fmt.Printf("===>%s \n", string(output))

	return nil
}
