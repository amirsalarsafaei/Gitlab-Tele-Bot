package merge

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strconv"
)

var diffRegex = regexp.MustCompile(`@@ -\d+,(\d+) \+\d+,(\d)+ @@`)

type DiffResponse []Diff

type Diff struct {
	Diff        string `json:"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

func (d Diff) GetDiffChanges() (int, int) {
	var removed, added int
	for _, change := range diffRegex.FindAllStringSubmatch(d.Diff, -1) {
		r, _ := strconv.Atoi(change[1])
		a, _ := strconv.Atoi(change[2])
		removed += r
		added += a
	}

	return removed, added
}

func (h Handler) getMergeRequestDiff(ctx context.Context, repoID int, mergeID int) ([]Diff, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("https://%s/api/v4/projects/%d/merge_requests/%d/diffs", h.config.Diff.GitLabHost, repoID, mergeID),
		nil,
	)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"repo_id":  repoID,
			"merge_id": mergeID,
		}).Error("could not create http request")
		return nil, errors.Wrap(err, "could not create http request")
	}

	req.Header.Set("PRIVATE-TOKEN", h.config.Diff.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"repo_id":  repoID,
			"merge_id": mergeID,
		}).Error("could send http request")
		return nil, errors.Wrap(err, "could send http request")
	}

	diffResponse := DiffResponse{}

	err = json.NewDecoder(resp.Body).Decode(&diffResponse)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"repo_id":  repoID,
			"merge_id": mergeID,
		}).Error("could decode diff response")
		return nil, errors.Wrap(err, "could not decode response")
	}
	return diffResponse, nil
}
