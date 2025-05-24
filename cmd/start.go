package cmd

import (
	"log/slog"

	"github.com/nabhdeep/gateway-cli/pkg/gateway"
	"github.com/spf13/cobra"
)

var daemonize bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the gateway server",
	Long:  `Starts the api gateway server with the loaded config. If config not loaded then API server starts with default configs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if daemonize {
			return runInBackground()
		}

		// Your normal server code
		gateway.Inti_API_gateway()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// Add the daemon flag to run in background
	startCmd.Flags().BoolVarP(&daemonize, "daemon", "d", false, "Run server in background")
}

func runInBackground() error {
	slog.Error("TODO")
	return nil
}
