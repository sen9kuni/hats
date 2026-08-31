package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var profileRemoveCmd = &cobra.Command{
	Use:   "remove [profile-name]",
	Short: "remove a git profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileID := args[0]
		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if _, exists := cfg.Profiles[profileID]; !exists {
			return fmt.Errorf("profile '%s' not found", profileID)
		}

		delete(cfg.Profiles, profileID)

		updatedRule := []config.Rule{}

		for _, r := range cfg.Rules {
			if r.Profile == profileID {
				continue
			}
			updatedRule = append(updatedRule, r)
		}

		cfg.Rules = updatedRule

		if err := config.Save(cfg); err != nil {
			return err
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("profile deleted, but failed to update to git: %w", err)
		}

		fmt.Printf("Removeed profile: %s\n", profileID)
		return nil
	},
}

func init() {
	profileCmd.AddCommand(profileRemoveCmd)
}
