package merge

type DiffConfig struct {
	Enabled    bool   `json:"enabled"`
	APIKey     string `json:"apiKey"`
	GitLabHost string `json:"gitlabHost"`
}

type MergeConfig struct {
	Diff DiffConfig `json:"diff"`
}
