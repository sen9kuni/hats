package cmd

import "github.com/spf13/cobra"

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remote URL mapping rules",
}

func init() {
	RootCmd.AddCommand(remoteCmd)
}
