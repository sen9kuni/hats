package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	gUtl "github.com/sen9kuni/hats/internal/utils"
)

type Profile struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type Config struct {
	Profiles map[string]Profile `toml:"profiles"`
}

func GetPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, gUtl.AppName, gUtl.ConfigFileName), nil
}

func Load() (*Config, error) {
	cfg := &Config{Profiles: make(map[string]Profile)}

	path, err := GetPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		if err := Save(cfg); err != nil {
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

	return cfg, nil
}

func Save(cfg *Config) error {
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
