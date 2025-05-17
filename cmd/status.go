package cmd

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks the status of the gateway server",
	Long:  `Checks if the api gateway server is running in the background and displays its status.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return checkServerStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func checkServerStatus() error {
	return nil
}
