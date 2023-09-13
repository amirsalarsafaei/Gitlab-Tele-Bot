package serve

import (
	"fmt"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd/config"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients/telegram"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/merge"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"time"
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

	brokers := []clients.MessageBroker{
		telegram.NewBot(conf),
	}

	mergeHandler := merge.Handler{
		Brokers: brokers,
	}

	r := mux.NewRouter()
	r.HandleFunc("/merge", mergeHandler.Notifier)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
