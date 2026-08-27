package cmd

import (
	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "generate file config profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
