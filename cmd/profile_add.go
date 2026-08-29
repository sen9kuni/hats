package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var (
	profileName  string
	profileEmail string
	profileKey   string
)

var profileAddCmd = &cobra.Command{
	Use:   "add [profile-name]",
	Short: "Add or update a Git Profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileID := args[0]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if cfg.Profiles == nil {
			cfg.Profiles = make(map[string]config.Profile)
		}

		cfg.Profiles[profileID] = config.Profile{
			Name:       profileName,
			Email:      profileEmail,
			SigningKey: profileKey,
		}

		if err := config.Save(cfg); err != nil {
			return err
		}

		fmt.Printf("Profile '%s' saved successfully. \n", profileID)
		return nil
	},
}

func init() {
	profileAddCmd.Flags().StringVarP(&profileName, "name", "n", "", "Git user.name (required)")
	profileAddCmd.Flags().StringVarP(&profileEmail, "email", "e", "", "Git user.email (required)")
	profileAddCmd.Flags().StringVarP(&profileName, "signing-key", "k", "", "Git user.SigningKey (optional)")

	profileAddCmd.MarkFlagRequired("name")
	profileAddCmd.MarkFlagRequired("email")

	profileCmd.AddCommand(profileAddCmd)
}
