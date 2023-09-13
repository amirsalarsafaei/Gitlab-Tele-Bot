package cmd

import (
	"github.com/spf13/cobra"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd/serve"
)

var (
	RootCmd = &cobra.Command{
		Use:   "notifier-cli",
		Short: "Notifier Command Line Interface",
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(serve.Cmd)
}
