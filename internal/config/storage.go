// Package config for handle Reads/writes ~/.config/hats/hats.toml
package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func GetPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "hats"), nil
}

func GenerateConfigFile() error {
	cfg := &Config{
		Profiles: make(map[string]Profile),
		Rules:    make([]Rule, 0),
	}
	return Save(cfg)
}

func Load() (*Config, error) {
	cfg := &Config{
		Profiles: make(map[string]Profile),
		Rules:    make([]Rule, 0),
	}

	path, err := GetPath()
	if err != nil {
		return nil, err
	}
	path = filepath.Join(path, "hats.toml")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		if err := GenerateConfigFile(); err != nil {
			return nil, err
		}
		return cfg, nil
	} else if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Profiles == nil {
		cfg.Profiles = make(map[string]Profile)
	}

	if cfg.Rules == nil {
		cfg.Rules = make([]Rule, 0)
	}

	return cfg, nil
}

func Save(cfg *Config) error {
	path, err := GetPath()
	if err != nil {
		return err
	}
	path = filepath.Join(path, "hats.toml")

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
