package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all saved git profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("error loading config: %s", err)
		}

		if len(cfg.Profiles) == 0 {
			return fmt.Errorf("no profiles found")
		}

		fmt.Println("Saved Profiles:")
		for key, profile := range cfg.Profiles {
			fmt.Printf("- %s:\n Name: %s\n Email: %s\n", key, profile.User.Name, profile.User.Email)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
