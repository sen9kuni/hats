package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version    = "dev"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hats version %s\n", Version)
		},
	}
)

func init() {
	RootCmd.AddCommand(versionCmd)
}
