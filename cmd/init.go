package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "append out hook config to main git config",
	RunE: func(cmd *cobra.Command, args []string) error {
		expectedPath := fmt.Sprintf("~/.config/hats/%s", config.IncludesFileName)
		if err := git.EnsureHookExists(expectedPath); err != nil {
			return err
		}

		fmt.Println("sucess to append hook config")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
