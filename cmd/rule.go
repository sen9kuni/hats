package cmd

import "github.com/spf13/cobra"

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Managing directory rules for Git Profile",
}

func init() {
	RootCmd.AddCommand(ruleCmd)
}
