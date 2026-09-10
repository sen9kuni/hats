package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var (
	updateProfileName       string
	updateProfileEmail      string
	updateProfileSigningKey string
	updateProfileSSHKey     string
)

var profileUpdateCmd = &cobra.Command{
	Use:   "update [profile-name]",
	Short: "Update an existing Git profile",
	Example: `	hats profile update work --email "new-email@work.com"
	hats profile update work -n "new name" -k "new-ssh-key"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileID := args[0]
		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		profile, exists := cfg.Profiles[profileID]
		if !exists {
			return fmt.Errorf("profile '%s' does not exist", profileID)
		}

		updated := false

		if cmd.Flags().Changed("name") {
			profile.Name = updateProfileName
			updated = true
		}
		if cmd.Flags().Changed("email") {
			profile.Email = updateProfileEmail
			updated = true
		}
		if cmd.Flags().Changed("signing-key") {
			profile.SigningKey = updateProfileSigningKey
			updated = true
		}
		if cmd.Flags().Changed("ssh-key") {
			profile.SSHKey = updateProfileSSHKey
			updated = true
		}

		if !updated {
			fmt.Println("No updated flags provided. Profile remain unchanged.")
			return nil
		}

		cfg.Profiles[profileID] = profile

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("profile updated, but failed to update to git: %w", err)
		}

		fmt.Printf("Profile '%s' updated successfully.\n", profileID)

		return nil
	},
}

func init() {
	profileCmd.AddCommand(profileUpdateCmd)

	profileUpdateCmd.Flags().StringVarP(&updateProfileName, "name", "n", "", "Update Git user.name")
	profileUpdateCmd.Flags().StringVarP(&updateProfileEmail, "email", "e", "", "Update Git user.email")
	profileUpdateCmd.Flags().StringVarP(&updateProfileSigningKey, "signing-key", "k", "", "Update Git user.signingKey")
	profileUpdateCmd.Flags().StringVarP(&updateProfileSSHKey, "ssh-key", "s", "", "Update Path to SSH key Git")
}
