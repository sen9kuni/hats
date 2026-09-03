package config

type Config struct {
	Profiles map[string]Profile `toml:"profiles"`
	Rules    []Rule             `toml:"rules"`
}

type Profile struct {
	Name       string `toml:"name"`
	Email      string `toml:"email"`
	SigningKey string `toml:"signing_key,omitempty"`
}

type Rule struct {
	Profile string `toml:"profile"`
	Path    string `toml:"path"`
}
