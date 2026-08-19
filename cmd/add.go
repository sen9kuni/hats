package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		cfg, err := config.Load()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		cfg.Profiles[key] = config.Profile{
			User: config.User{
				Name:  nameFlag,
				Email: emailFlag,
			},
		}

		if err := config.Save(cfg); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}
		fmt.Printf("Profile '%s' save successfully!\n", key)
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
