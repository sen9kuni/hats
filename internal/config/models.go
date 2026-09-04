package config

type Config struct {
	Profiles    map[string]Profile `toml:"profiles"`
	Rules       []Rule             `toml:"rules"`
	RemoteRules []Remote           `toml:"remote_rules"`
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

type Remote struct {
	URL     string `toml:"url"`
	Profile string `toml:"profile"`
}
