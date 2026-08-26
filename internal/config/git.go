package config

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	gUtl "github.com/sen9kuni/hats/internal/utils"
)

type GitConfig struct {
	Profiles map[string]map[string]map[string]string
}

func GetPathGitFile() (string, error) {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, gUtl.GitFileName), nil
}

func GetListGitConfig() (*GitConfig, error) {
	cfg := &GitConfig{Profiles: make(map[string]map[string]map[string]string)}
	mainConfig, err := ScannerGit("")
	if err != nil {
		return nil, err
	}
	cfg.Profiles["main"] = mainConfig

	if len(cfg.Profiles) > 0 && cfg.Profiles["main"] != nil && cfg.Profiles["main"]["includeif"] != nil && len(cfg.Profiles["main"]["includeif"]) > 0 {
		for KeyIncludeif, PIncludeif := range cfg.Profiles["main"]["includeif"] {
			newCfg, err := ScannerGit(PIncludeif)
			if err != nil {
				return nil, err
			}
			cfg.Profiles[KeyIncludeif] = newCfg
		}
	}
	return cfg, nil
}

func ScannerGit(pathConfig string) (map[string]map[string]string, error) {
	result := make(map[string]map[string]string)
	var cmd *exec.Cmd
	if len(pathConfig) > 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(pathConfig, "~/") {
			pathConfig = filepath.Join(homeDir, pathConfig[2:])
		} else if pathConfig == "~" {
			pathConfig = homeDir
		}
		cmd = exec.Command("git", "config", "--file", pathConfig, "--list")
	} else {
		cmd = exec.Command("git", "config", "--global", "--list")
	}

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	for line := range strings.SplitSeq(string(output), "\n") {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		key := parts[0]
		value := ""

		if len(parts) == 2 {
			value = parts[1]
		}

		key = strings.ToLower(key)
		keyParts := strings.SplitN(key, ".", 2)
		mainKey := keyParts[0]
		subKey := ""
		if len(keyParts) == 2 {
			subKey = keyParts[1]
		}

		if result[mainKey] == nil {
			result[mainKey] = make(map[string]string)
		}

		result[mainKey][subKey] = value
	}

	return result, nil
}

func SyncToConfig(cfg *GitConfig) error {
	path, err := GetPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func ApplyConfig() error {
	return nil
}
