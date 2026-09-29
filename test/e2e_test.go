package test_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sen9kuni/hats/cmd"
)

func runCLI(t *testing.T, args ...string) {
	t.Helper()

	c := cmd.RootCmd

	c.SetArgs(args)

	var out bytes.Buffer
	c.SetOut(&out)
	c.SetErr(&out)

	err := c.Execute()
	if err != nil {
		t.Fatalf("Command 'hats %v' failed: %v\nOutput: %s", args, err, out.String())
	}
}

func TestHatsEndToEnd(t *testing.T) {
	mockHome := t.TempDir()

	t.Setenv("Home", mockHome)
	t.Setenv("USERPROFILE", mockHome)

	workDir := filepath.Join(mockHome, "work")
	os.MkdirAll(workDir, 0o775)

	runCLI(t, "init")

	runCLI(t, "profile", "add", "work", "--name", "Work Jane", "--email", "jane@work.com")

	runCLI(t, "rule", "add", "work", filepath.Join(mockHome, "work"))

	projectDir := filepath.Join(workDir, "project-a")
	os.MkdirAll(projectDir, 0o755)

	gitInit := exec.Command("git", "init")
	gitInit.Dir = projectDir
	if err := gitInit.Run(); err != nil {
		t.Fatalf("failed to git init: %v", err)
	}

	gitConfig := exec.Command("git", "config", "--get", "user.email")
	gitConfig.Dir = projectDir

	var out bytes.Buffer
	gitConfig.Stdout = &out
	if err := gitConfig.Run(); err != nil {
		t.Fatalf("git config failed (did includeIf fail?): %v", err)
	}

	actualEmail := strings.TrimSpace(out.String())
	expectedEmail := "jane@work.com"

	if actualEmail != expectedEmail {
		t.Errorf("expected email %q, but Git resolved %q", expectedEmail, actualEmail)
	}
}
