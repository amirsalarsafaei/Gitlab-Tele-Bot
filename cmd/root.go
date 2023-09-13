package cmd

import (
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd/serve"
	"github.com/spf13/cobra"
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
