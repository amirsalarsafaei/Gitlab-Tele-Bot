package git

type GitConfig struct {
	APIKey     string `mapstructure:"apiKey"`
	GitLabHost string `mapstructure:"gitlabHost"`
}
