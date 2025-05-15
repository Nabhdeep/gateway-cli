package cmd

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the gateway server",
	Long:  `Starts the api gateway server with the loaded config. If config not loaded then API server starts with default configs.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
