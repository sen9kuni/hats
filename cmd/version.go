package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version    = "dev"
	Commit     = "none"
	Date       = "unknown"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hats v%s\n", Version)
			fmt.Printf("Commit: %s\n", Commit)
			fmt.Printf("Built: %s\n", Date)
		},
	}
)

func init() {
	RootCmd.AddCommand(versionCmd)
}
