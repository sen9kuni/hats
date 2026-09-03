// Package git for Wrapper for exec.Command("git", ...)
package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func EnsureHookExists(expectedPath string) error {
	cmd := exec.Command("git", "config", "--global", "--get-all", "include.path")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil && cmd.ProcessState.ExitCode() != 1 {
		return fmt.Errorf("failed to read global git config: %w", err)
	}

	lines := strings.SplitSeq(strings.TrimSpace(out.String()), "\n")
	for line := range lines {
		if line == expectedPath {
			return nil
		}
	}

	addCmd := exec.Command("git", "config", "--global", "--add", "include.path", expectedPath)
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to inject hats hook into global git config: %w", err)
	}
	return nil
}
