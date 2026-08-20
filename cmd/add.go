package cmd

import (
	"fmt"
	"strings"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var (
	nameFlag  string
	emailFlag string
)

var addCmd = &cobra.Command{
	Use:   "add [profile-key]",
	Short: "Add or Update a git profile",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !strings.Contains(emailFlag, "@") {
			return fmt.Errorf("invalid email: %s", emailFlag)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("error loading config: '%s'", err)
		}

		cfg.Profiles[key] = config.Profile{
			User: config.User{
				Name:  nameFlag,
				Email: emailFlag,
			},
		}

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("error saving config: '%s'", err)
		}
		fmt.Printf("Profile '%s' save successfully!\n", key)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&nameFlag, "name", "n", "", "Git user Name (required)")
	addCmd.Flags().StringVarP(&emailFlag, "email", "e", "", "Git user Email (required)")
	if err := addCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
	if err := addCmd.MarkFlagRequired("email"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(addCmd)
}
