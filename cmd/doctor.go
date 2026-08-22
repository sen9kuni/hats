package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "scan .gitconfig and store it on out data management",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ScanGitFile("", "")
		if err != nil {
			return err
		}

		if len(cfg.Profiles) == 0 {
			return fmt.Errorf("no profiles found")
		}

		fmt.Println("Scanned Profiles:")
		for key, profile := range cfg.Profiles {
			fmt.Printf("- %s:\n Name: %s\n Email: %s\n", key, profile["user"]["name"], profile["user"]["email"])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
