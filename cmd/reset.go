package cmd

import (
	"fmt"
	"strings"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset config git profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("This will delete all profile. Continue? (y/N): ")

		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Reset cancelled (input error).")
			return nil
		}

		if strings.ToLower(response) != "y" {
			fmt.Println("Reset cancelled.")
			return nil
		}

		cfg := &config.Config{
			Profiles: make(map[string]config.Profile),
		}

		if err := config.Save(cfg); err != nil {
			fmt.Println("Error reset profiles:", err)
			return nil
		}
		fmt.Println("All profiles have been reset.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
