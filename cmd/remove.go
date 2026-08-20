package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [profile-key]",
	Short: "remove a git profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("error loading config: %s", err)
		}

		if _, exists := cfg.Profiles[key]; !exists {
			return fmt.Errorf("profile '%s' does not exist", key)
		}

		delete(cfg.Profiles, key)

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("error saving config: %s", err)
		}

		fmt.Printf("Profile '%s' deleted successfully.\n", key)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
