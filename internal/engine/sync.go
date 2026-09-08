package engine

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sen9kuni/hats/internal/config"
)

func Sync(cfg *config.Config, hatsDir string) error {
	profileActiveDir := filepath.Join(hatsDir, "profiles")
	profileTmpDir := filepath.Join(hatsDir, "profiles_tmp")
	includesActiveFile := filepath.Join(hatsDir, config.IncludesFileName)
	includesTmpFile := filepath.Join(hatsDir, config.IncludesFileNameTemp)

	if err := os.RemoveAll(profileTmpDir); err != nil {
		return fmt.Errorf("failed to remove temp profiles: %w", err)
	}

	if err := os.MkdirAll(profileTmpDir, 0o755); err != nil {
		return fmt.Errorf("fail to create temp profiles dir: %s", err)
	}

	for name, profile := range cfg.Profiles {
		content := GenerateProfileConfig(profile)
		tmpPath := filepath.Join(profileTmpDir, name+".gitconfig")

		if err := os.WriteFile(tmpPath, []byte(content), 0o664); err != nil {
			return fmt.Errorf("failed to write profile %s: %w", name, err)
		}
	}

	includesContent := GenerateIncludesConfig(cfg.Rules, profileActiveDir)
	if err := os.WriteFile(includesTmpFile, []byte(includesContent), 0o644); err != nil {
		return fmt.Errorf("failed to write includes: %w", err)
	}

	includesRemoteContent := GenerateIncludesRemoteConfig(cfg.RemoteRules, profileActiveDir)
	if err := os.WriteFile(includesRemoteContent, []byte(includesContent), 0o644); err != nil {
		return fmt.Errorf("failed to write remotes: %w", err)
	}

	if err := os.RemoveAll(profileActiveDir); err != nil {
		return fmt.Errorf("failed to remove active profiles: %w", err)
	}

	if err := os.Rename(profileTmpDir, profileActiveDir); err != nil {
		return fmt.Errorf("failed to active profiles: %w", err)
	}

	if err := os.Rename(includesTmpFile, includesActiveFile); err != nil {
		return fmt.Errorf("fail to active includes: %w", err)
	}
	return nil
}
