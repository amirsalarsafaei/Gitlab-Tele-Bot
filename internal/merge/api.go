package merge

import "time"

type APIModel struct {
	Id          int       `json:"id"`
	Iid         int       `json:"iid"`
	ProjectId   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	MergedBy    struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		Locked    bool   `json:"locked"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"merged_by"`
	MergeUser struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		Locked    bool   `json:"locked"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"merge_user"`
	MergedAt                  time.Time   `json:"merged_at"`
	ClosedBy                  interface{} `json:"closed_by"`
	ClosedAt                  interface{} `json:"closed_at"`
	TargetBranch              string      `json:"target_branch"`
	SourceBranch              string      `json:"source_branch"`
	UserNotesCount            int         `json:"user_notes_count"`
	Upvotes                   int         `json:"upvotes"`
	Downvotes                 int         `json:"downvotes"`
	Author                    `json:"author"`
	Assignees                 []interface{} `json:"assignees"`
	Assignee                  interface{}   `json:"assignee"`
	Reviewers                 Reviewers     `json:"reviewers"`
	SourceProjectId           int           `json:"source_project_id"`
	TargetProjectId           int           `json:"target_project_id"`
	Labels                    []interface{} `json:"labels"`
	Draft                     bool          `json:"draft"`
	WorkInProgress            bool          `json:"work_in_progress"`
	Milestone                 interface{}   `json:"milestone"`
	MergeWhenPipelineSucceeds bool          `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string        `json:"merge_status"`
	DetailedMergeStatus       string        `json:"detailed_merge_status"`
	Sha                       string        `json:"sha"`
	MergeCommitSha            string        `json:"merge_commit_sha"`
	SquashCommitSha           string        `json:"squash_commit_sha"`
	DiscussionLocked          interface{}   `json:"discussion_locked"`
	ShouldRemoveSourceBranch  bool          `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool          `json:"force_remove_source_branch"`
	PreparedAt                time.Time     `json:"prepared_at"`
	Reference                 string        `json:"reference"`
	References                struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	WebUrl    string `json:"web_url"`
	TimeStats struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	SquashOnMerge        bool `json:"squash_on_merge"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool        `json:"has_conflicts"`
	BlockingDiscussionsResolved bool        `json:"blocking_discussions_resolved"`
	Subscribed                  bool        `json:"subscribed"`
	ChangesCount                string      `json:"changes_count"`
	LatestBuildStartedAt        time.Time   `json:"latest_build_started_at"`
	LatestBuildFinishedAt       time.Time   `json:"latest_build_finished_at"`
	FirstDeployedToProductionAt interface{} `json:"first_deployed_to_production_at"`
	Pipeline                    struct {
		Id        int       `json:"id"`
		Iid       int       `json:"iid"`
		ProjectId int       `json:"project_id"`
		Sha       string    `json:"sha"`
		Ref       string    `json:"ref"`
		Status    string    `json:"status"`
		Source    string    `json:"source"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		WebUrl    string    `json:"web_url"`
	} `json:"pipeline"`
	HeadPipeline struct {
		Id         int         `json:"id"`
		Iid        int         `json:"iid"`
		ProjectId  int         `json:"project_id"`
		Sha        string      `json:"sha"`
		Ref        string      `json:"ref"`
		Status     string      `json:"status"`
		Source     string      `json:"source"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		WebUrl     string      `json:"web_url"`
		BeforeSha  string      `json:"before_sha"`
		Tag        bool        `json:"tag"`
		YamlErrors interface{} `json:"yaml_errors"`
		User       struct {
			Id        int    `json:"id"`
			Username  string `json:"username"`
			Name      string `json:"name"`
			State     string `json:"state"`
			Locked    bool   `json:"locked"`
			AvatarUrl string `json:"avatar_url"`
			WebUrl    string `json:"web_url"`
		} `json:"user"`
		StartedAt      time.Time   `json:"started_at"`
		FinishedAt     time.Time   `json:"finished_at"`
		CommittedAt    interface{} `json:"committed_at"`
		Duration       int         `json:"duration"`
		QueuedDuration int         `json:"queued_duration"`
		Coverage       interface{} `json:"coverage"`
		DetailedStatus struct {
			Icon         string      `json:"icon"`
			Text         string      `json:"text"`
			Label        string      `json:"label"`
			Group        string      `json:"group"`
			Tooltip      string      `json:"tooltip"`
			HasDetails   bool        `json:"has_details"`
			DetailsPath  string      `json:"details_path"`
			Illustration interface{} `json:"illustration"`
			Favicon      string      `json:"favicon"`
		} `json:"detailed_status"`
	} `json:"head_pipeline"`
	DiffRefs struct {
		BaseSha  string `json:"base_sha"`
		HeadSha  string `json:"head_sha"`
		StartSha string `json:"start_sha"`
	} `json:"diff_refs"`
	MergeError        interface{} `json:"merge_error"`
	FirstContribution bool        `json:"first_contribution"`
	User              struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
}
