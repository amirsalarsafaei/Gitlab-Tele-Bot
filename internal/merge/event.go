package merge

import (
	"time"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/pkg/gitlab"
)

type Event struct {
	EventType  string `json:"event_type"`
	ObjectKind string `json:"object_kind"`

	Attributes struct {
		AssigneeId interface {
		} `json:"assignee_id"`
		AssigneeIds []interface {
		} `json:"assignee_ids"`
		AuthorId                    int       `json:"author_id"`
		BlockingDiscussionsResolved bool      `json:"blocking_discussions_resolved"`
		CreatedAt                   time.Time `json:"created_at"`
		Description                 string    `json:"description"`
		HeadPipelineId              int       `json:"head_pipeline_id"`
		HumanTimeChange             interface {
		} `json:"human_time_change"`
		HumanTimeEstimate interface {
		} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface {
		} `json:"human_total_time_spent"`
		Id     int `json:"id"`
		Iid    int `json:"iid"`
		Labels []interface {
		} `json:"labels"`
		LastCommit struct {
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Id        string    `json:"id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Title     string    `json:"title"`
			Url       string    `json:"url"`
		} `json:"last_commit"`
		LastEditedAt interface {
		} `json:"last_edited_at"`
		LastEditedById interface {
		} `json:"last_edited_by_id"`
		MergeCommitSha string `json:"merge_commit_sha"`
		MergeError     interface {
		} `json:"merge_error"`
		MergeParams struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus string `json:"merge_status"`
		MergeUserId interface {
		} `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool `json:"merge_when_pipeline_succeeds"`
		MilestoneId               interface {
		} `json:"milestone_id"`
		Source struct {
			AvatarUrl interface {
			} `json:"avatar_url"`
			CiConfigPath      string `json:"ci_config_path"`
			DefaultBranch     string `json:"default_branch"`
			Description       string `json:"description"`
			GitHttpUrl        string `json:"git_http_url"`
			GitSshUrl         string `json:"git_ssh_url"`
			Homepage          string `json:"homepage"`
			HttpUrl           string `json:"http_url"`
			Id                int    `json:"id"`
			Name              string `json:"name"`
			Namespace         string `json:"namespace"`
			PathWithNamespace string `json:"path_with_namespace"`
			SshUrl            string `json:"ssh_url"`
			Url               string `json:"url"`
			VisibilityLevel   int    `json:"visibility_level"`
			WebUrl            string `json:"web_url"`
		} `json:"source"`
		SourceBranch    string `json:"source_branch"`
		SourceProjectId int    `json:"source_project_id"`
		State           string `json:"state"`
		StateId         int    `json:"state_id"`
		Target          struct {
			AvatarUrl interface {
			} `json:"avatar_url"`
			CiConfigPath      string `json:"ci_config_path"`
			DefaultBranch     string `json:"default_branch"`
			Description       string `json:"description"`
			GitHttpUrl        string `json:"git_http_url"`
			GitSshUrl         string `json:"git_ssh_url"`
			Homepage          string `json:"homepage"`
			HttpUrl           string `json:"http_url"`
			Id                int    `json:"id"`
			Name              string `json:"name"`
			Namespace         string `json:"namespace"`
			PathWithNamespace string `json:"path_with_namespace"`
			SshUrl            string `json:"ssh_url"`
			Url               string `json:"url"`
			VisibilityLevel   int    `json:"visibility_level"`
			WebUrl            string `json:"web_url"`
		} `json:"target"`
		TargetBranch    string    `json:"target_branch"`
		TargetProjectId int       `json:"target_project_id"`
		TimeChange      int       `json:"time_change"`
		TimeEstimate    int       `json:"time_estimate"`
		Title           string    `json:"title"`
		TotalTimeSpent  int       `json:"total_time_spent"`
		UpdatedAt       time.Time `json:"updated_at"`
		UpdatedById     interface {
		} `json:"updated_by_id"`
		Url            string `json:"url"`
		WorkInProgress bool   `json:"work_in_progress"`
	} `json:"object_attributes"`

	gitlab.User

	gitlab.Project
	gitlab.Repository
}
