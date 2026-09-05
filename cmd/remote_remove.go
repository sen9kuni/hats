package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var remoteRemoveCmd = &cobra.Command{
	Use:   "remove [url-pattern]",
	Short: "remove a Git remote URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		urlPattern := args[0]
		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		updatedRemote := []config.Remote{}
		removed := false

		for _, r := range cfg.RemoteRules {
			if r.URL == urlPattern {
				removed = true
				continue
			}
			updatedRemote = append(updatedRemote, r)
		}

		if !removed {
			return fmt.Errorf("no remote found for url pattern: %s", urlPattern)
		}

		cfg.RemoteRules = updatedRemote

		err = config.Save(cfg)
		if err != nil {
			return err
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("remote deleted, but failed to apply to git: %w", err)
		}

		fmt.Printf("removed remote for: %s\n", urlPattern)

		return nil
	},
}

func init() {
	remoteCmd.AddCommand(remoteRemoveCmd)
}
