package run

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd/config"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients/stdout"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients/telegram"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/merge"
)

var (
	Cmd = &cobra.Command{
		Use:   "run",
		Short: "runs merge request handler",
		RunE:  run,
		Args:  cobra.ExactArgs(2),
	}
)

func run(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) != 2 {
		return errors.New("command must be run with [project id] [merge request iid]")
	}

	projectID, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("project id must be int")
	}
	mergeID, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("merge id must be int")
	}

	conf := config.LoadConfig()

	telegramBot, err := telegram.NewBot(conf.Telegram)
	if err != nil {
		return err
	}

	brokers := []clients.MessageBroker{
		telegramBot,
		stdout.NewPrintBot(),
	}

	mergeHandler := merge.NewHandler(brokers, conf.Merge, conf.Git)

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("https://%s/api/v4/projects/%d/merge_requests/%d/",
			conf.Git.GitLabHost, projectID, mergeID),
		nil,
	)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"repo_id":  projectID,
			"merge_id": mergeID,
		}).Error("could not create http request")
		return errors.Wrap(err, "could not create http request")
	}

	req.Header.Set("PRIVATE-TOKEN", conf.Git.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"repo_id":  projectID,
			"merge_id": mergeID,
		}).Error("could send http request")
		return errors.Wrap(err, "could send http request")
	}

	mergeResp := merge.APIModel{}
	err = json.NewDecoder(resp.Body).Decode(&mergeResp)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"repo_id":  projectID,
			"merge_id": mergeID,
		}).Error("could decode diff response")
		return errors.Wrap(err, "could not decode response")
	}
	return mergeHandler.Handle(ctx, projectID, mergeID, mergeResp.Title, mergeResp.Description, mergeResp.WebUrl,
		mergeResp.Author, mergeResp.Reviewers)
}
