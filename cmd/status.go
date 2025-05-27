package cmd

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks the status of the gateway server",
	Long:  `Checks if the api gateway server is running in the background and displays its status.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			return errors.New("invalid flags")
		}

		pid, err := cmd.Flags().GetInt("pid")
		if err != nil {
			return errors.New("invalid process id")
		}
		if pid != 0 {
			return checkServerStatus(pid)
		}

		return checkServerStatus(0)
	},
}

func init() {
	statusCmd.Flags().IntP("pid", "p", 0, "Process ID of the gateway server to check")
	rootCmd.AddCommand(statusCmd)
}

// https://stackoverflow.com/questions/15204162/check-if-a-process-exists-in-go-way
func checkServerStatus(pid int) error {
	if pid == 0 {
		pidd, err := GetDemonProcessID()
		if err != nil {
			return err
		}
		pid = pidd
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	} else {
		err := process.Signal(syscall.Signal(0))
		switch {
		case err == nil:
			fmt.Printf("✓ Gateway server (PID %d) is running\n", pid)
			return nil
		case errors.Is(err, os.ErrProcessDone):
			fmt.Printf("✗ Gateway server process has finished\n")
			return fmt.Errorf("gateway server is not running")
		default:
			fmt.Printf("✗ Gateway server (PID %d) is not accessible: %v\n", pid, err)
			return fmt.Errorf("gateway server status unknown: %w", err)
		}
	}
}
