package cmd

import (
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

func stopBackgroundServer() error {
	return nil
}
