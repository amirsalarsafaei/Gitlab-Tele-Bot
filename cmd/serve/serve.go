package serve

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd/config"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients/stdout"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients/telegram"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/merge"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/pkg/gitlab"
)

var (
	Cmd = &cobra.Command{
		Use:   "serve",
		Short: "serves gitlab notifier bots",
		RunE:  serve,
		Args:  cobra.ExactArgs(1),
	}
)

func serve(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	log.WithContext(ctx).Info("Starting Serve Command")

	port := args[0]
	conf := config.LoadConfig()

	telegramBot, err := telegram.NewBot(conf.Telegram)
	if err != nil {
		return err
	}

	brokers := []clients.MessageBroker{
		telegramBot,
		stdout.NewPrintBot(),
	}

	mergeHandler := merge.Handler{
		Brokers: brokers,
	}

	r := mux.NewRouter()
	r.Use(gitlab.NewGitLabAuthMiddleWare(conf.Secret))

	r.HandleFunc("/merge", mergeHandler.Notifier)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
