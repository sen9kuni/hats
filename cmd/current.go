package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/git"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "show the active Git profile for the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		activeEmail := git.GetValue("user.email")

		if activeEmail == "" {
			fmt.Println("No Git email is currently configured for this directory")
			return nil
		}

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load hats config: %w", err)
		}

		var matchedProfileName string
		for name, profile := range cfg.Profiles {
			if profile.Email == activeEmail {
				matchedProfileName = name
			}
		}

		if matchedProfileName != "" {
			fmt.Printf("	Active Profile: %s\n", matchedProfileName)
			fmt.Printf("	Email: %s\n", activeEmail)
		} else {
			fmt.Printf("	Active Identity: Global Default\n")
			fmt.Printf("	Email: %s\n", activeEmail)
			fmt.Printf("	(This email does not match any managed Hats Profile)\n")
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(currentCmd)
}
