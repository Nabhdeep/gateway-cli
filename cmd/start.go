package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"syscall"

	"github.com/nabhdeep/gateway-cli/pkg/constants"
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
	processID, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if err != 0 {
		slog.Error("Error in creatig a syscall to Fork")
		return errors.New("error in fork")
	}

	if processID > 0 {
		// killing the parent process
		os.Exit(0)
	}
	pid, errr := syscall.Setsid()
	if errr != nil {
		return errr
	}
	file, r := os.Create(constants.Gateway_Pid)
	if r != nil {
		return r
	}
	defer file.Close()

	_, erro := file.WriteString(fmt.Sprintf("%d", pid))
	if erro != nil {
		slog.Error("Unable to create a pid")
		return erro
	}

	// starting the process
	gateway.Inti_API_gateway()

	return nil
}
