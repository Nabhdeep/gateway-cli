package cmd

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/nabhdeep/gateway-cli/pkg/constants"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the gateway server",
	Long:  `Stops the api gateway server that is running in the background.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return stopBackgroundServer()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func GetDemonProcessID() (int, error) {
	f, err := os.ReadFile(constants.Gateway_Pid)
	if err != nil {
		slog.Error("unable to read gatway.pid file or its missing")
		return 0, err
	}
	pidInt, err := strconv.Atoi(string(f))
	if err != nil {
		return 0, err
	}
	return pidInt, nil
}

func stopBackgroundServer() error {
	defer os.Remove(constants.Gateway_Pid)
	pid, err := GetDemonProcessID()
	if err != nil {
		return err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		slog.Error("unbale to find process with ", slog.Int("pid", pid))
		return err
	}

	err = process.Kill()
	if err != nil {
		slog.Error("unable to kil process with ", slog.Int("pid", pid))
		return err
	}
	return nil
}
