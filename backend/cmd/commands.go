package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "configurator cloud manage resources.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(runCmd)
}
