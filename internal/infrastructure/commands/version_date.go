package commands

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type ModuleInfo struct {
	Time string `json:"Time"`
}

func fetchModuleVersionDate(module, version string) (time.Time, error) {
	cmd := exec.Command("go", "list", "-m", "-json", fmt.Sprintf("%s@%s", module, version))
	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, fmt.Errorf("error executing command: %w", err)
	}

	var info ModuleInfo
	decoder := json.NewDecoder(strings.NewReader(string(output)))
	if err := decoder.Decode(&info); err != nil {
		return time.Time{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	pubDate, err := time.Parse(time.RFC3339, info.Time)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %w", err)
	}

	return pubDate, nil
}

func IsSafityModuleVersion(module, version string, timeThreshold time.Duration) (bool, error) {
	pubDate, err := fetchModuleVersionDate(module, version)
	if err != nil {
		return false, err
	}

	elapsed := time.Since(pubDate)
	if elapsed >= timeThreshold {
		return true, nil
	}

	return false, nil
}
