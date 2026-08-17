package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all saved git profiles",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		if len(cfg.Profiles) == 0 {
			fmt.Println("No profiles found.")
			return
		}

		fmt.Println("Saved Profiles:")
		for key, profile := range cfg.Profiles {
			fmt.Printf("- %s:\n Name: %s\n Email: %s\n", key, profile.Name, profile.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
