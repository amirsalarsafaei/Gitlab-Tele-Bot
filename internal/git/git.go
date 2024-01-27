package git

type GitConfig struct {
	APIKey     string `mapstructure:"apikey"`
	GitLabHost string `mapstructure:"gitlabhost"`
}
