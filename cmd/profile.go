package cmd

import "github.com/spf13/cobra"

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Managing Git Profile",
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
