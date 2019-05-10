package config

type Command struct {
	Description string   `mapstructure:"description"`
	Workers     int      `mapstructure:"workers"`
	Onerror     string   `mapstructure:"onerror"`
	Name        string   `mapstructure:"name"`
	Args        []string `mapstructure:"args"`
}

type Config struct {
	Commit   string              `mapstructure:"commit"`
	WalkPath string              `mapstructure:"walk-path"`
	Commands map[string]*Command `mapstructure:"commands"`
}
