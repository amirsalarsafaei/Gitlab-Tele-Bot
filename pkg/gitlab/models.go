package gitlab

type User struct {
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
}

type Repository struct {
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Name        string `json:"name"`
	Url         string `json:"url"`
}

type Project struct {
	AvatarUrl         interface{} `json:"avatar_url"`
	CiConfigPath      string      `json:"ci_config_path"`
	DefaultBranch     string      `json:"default_branch"`
	Description       string      `json:"description"`
	GitHttpUrl        string      `json:"git_http_url"`
	GitSshUrl         string      `json:"git_ssh_url"`
	Homepage          string      `json:"homepage"`
	HttpUrl           string      `json:"http_url"`
	Id                int         `json:"id"`
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	PathWithNamespace string      `json:"path_with_namespace"`
	SshUrl            string      `json:"ssh_url"`
	Url               string      `json:"url"`
	VisibilityLevel   int         `json:"visibility_level"`
	WebUrl            string      `json:"web_url"`
}
