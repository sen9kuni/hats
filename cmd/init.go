package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/sen9kuni/hats/internal/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Install the Hats include hook into global git config",
	RunE: func(cmd *cobra.Command, args []string) error {
		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		expectedPath := filepath.Join(hatsDir, config.IncludesFileName)
		expectedPath = engine.FormatToTilde(expectedPath)

		if err := git.EnsureHookExists(expectedPath); err != nil {
			return err
		}

		fmt.Println("success to append hook config")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
