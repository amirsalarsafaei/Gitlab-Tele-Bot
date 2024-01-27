package merge

type DiffConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type MergeConfig struct {
	Diff DiffConfig `mapstructure:"diff"`
}
