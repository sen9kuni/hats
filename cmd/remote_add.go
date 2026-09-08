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

		for _, r := range cfg.RemoteRules {
			if r.URL == urlPattern {
				return fmt.Errorf("remote rule for '%s' already exists", urlPattern)
			}
		}

		cfg.RemoteRules = append(cfg.RemoteRules, config.Remote{
			URL:     urlPattern,
			Profile: profileID,
		})

		if err := config.Save(cfg); err != nil {
			return err
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("remote saved, but failed to apply to git: %w", err)
		}

		fmt.Printf("Mapped remote URL '%s' to profile '%s'\n", urlPattern, profileID)
		return nil
	},
}

func init() {
	remoteCmd.AddCommand(remoteAddCmd)
}
