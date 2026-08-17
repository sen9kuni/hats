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
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		cfg, err := config.Load()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		if _, exists := cfg.Profiles[key]; !exists {
			fmt.Printf("Profile '%s' does not exist.\n", key)
			return
		}

		delete(cfg.Profiles, key)

		if err := config.Save(cfg); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

		fmt.Printf("Profile '%s' deleted successfully.\n", key)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
