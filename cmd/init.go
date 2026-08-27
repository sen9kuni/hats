package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "append out hook config to main git config",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := git.EnsureHookExists("~/.config/hats/managed_includes.gitconfig"); err != nil {
			return err
		}

		fmt.Println("sucess to append hook config")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
