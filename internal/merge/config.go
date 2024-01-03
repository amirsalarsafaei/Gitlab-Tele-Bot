package merge

type DiffConfig struct {
	Enabled    bool   `json:"enabled"`
	APIKey     string `json:"api_key"`
	GitLabHost string `json:"gitlab_host"`
}

type MergeConfig struct {
	Diff DiffConfig `json:"diff"`
}
