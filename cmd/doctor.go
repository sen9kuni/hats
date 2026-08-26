package cmd

import (
	"fmt"
	"strings"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "scan .gitconfig and store it on out data management",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetListGitConfig()
		if err != nil {
			return err
		}

		if len(cfg.Profiles) == 0 {
			return fmt.Errorf("no profiles found")
		}

		fmt.Println("Scanned Current Git Profiles:")
		for key, profile := range cfg.Profiles {
			fmt.Printf("- %s:\n Name: %s\n Email: %s\n", key, profile["user"]["name"], profile["user"]["email"])
		}

		fmt.Printf("This will save current .gitconfig to storage config, Continue? (y/N): ")

		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Save cancelled (input error).")
			return nil
		}

		if strings.ToLower(response) != "y" {
			fmt.Println("Save cancelled.")
			return nil
		}
		if err := config.SyncToConfig(cfg); err != nil {
			fmt.Println("Error Save profiles:", err)
			return nil
		}

		fmt.Println("All git profiles have been save.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
