package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var remoteAddCmd = &cobra.Command{
	Use:   "add [url-pattern] [profile-id]",
	Short: "Map a profile to a Git remote URL",
	Example: `hats remote add "git@github.com:my-company/**" work
	hats remote add "https://gitlab.com:client-xyz/**" client1`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		urlPattern := args[0]
		profileID := args[1]

		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if _, exists := cfg.Profiles[profileID]; !exists {
			return fmt.Errorf("profile '%s' does not exist", profileID)
		}

		remoteUpdated := false
		for i, r := range cfg.RemoteRules {
			if r.URL == urlPattern {
				cfg.RemoteRules[i].Profile = profileID
				remoteUpdated = true
				break
			}
		}

		if !remoteUpdated {
			cfg.RemoteRules = append(cfg.RemoteRules, config.Remote{
				URL:     urlPattern,
				Profile: profileID,
			})
		}

		if err := config.Save(cfg); err != nil {
			return err
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("remote saved, but failed to apply to git: %w", err)
		}

		if remoteUpdated {
			fmt.Printf("Updated remote: %s now uses profile '%s'\n", urlPattern, profileID)
		} else {
			fmt.Printf("Mapped remote URL '%s' to profile '%s'\n", urlPattern, profileID)
		}
		return nil
	},
}

func init() {
	remoteCmd.AddCommand(remoteAddCmd)
}
