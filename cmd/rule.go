package cmd

import "github.com/spf13/cobra"

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Managing rule every Git Profile",
}

func init() {
	RootCmd.AddCommand(ruleCmd)
}
