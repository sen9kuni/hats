// Package cmd contains the application's Cobra CLI commands.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:               "Hats",
	Short:             "A Git profile manager",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
