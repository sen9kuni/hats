package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func ScanGitFile(nameProfile string, pathProfile string) (*GitConfig, error) {
	cfg := &GitConfig{Profiles: make(map[string]map[string]map[string]string)}
	var cmd *exec.Cmd
	profileName := "main"
	if len(nameProfile) > 0 && len(pathProfile) > 0 {
		cmd = exec.Command("git", "config", "--file", pathProfile, "--list")
		profileName = nameProfile
	} else {
		cmd = exec.Command("git", "config", "--global", "--list")
	}

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(output), "\n") {
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

		if cfg.Profiles[profileName] == nil {
			cfg.Profiles[profileName] = make(map[string]map[string]string)
		}

		if cfg.Profiles[profileName][mainKey] == nil {
			cfg.Profiles[profileName][mainKey] = make(map[string]string)
		}

		cfg.Profiles[profileName][mainKey][subKey] = value
	}

	if len(cfg.Profiles) > 0 && cfg.Profiles[profileName] != nil && cfg.Profiles["main"]["includeif"] != nil && len(cfg.Profiles["main"]["includeif"]) > 0 {
		// TODO: make recrusive for includeif
		fmt.Println("includeif have more than 0")
	}
	// fmt.Println(cfg)
	data := &cfg

	b, _ := json.MarshalIndent(*data, "", " ")
	fmt.Println(string(b))
	return cfg, nil
}

func ApplyConfig() error {
	return nil
}
