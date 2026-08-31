package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var ruleListCmd = &cobra.Command{
	Use:   "list",
	Short: "get list rule",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		_, err = fmt.Fprintln(w, "PROFILE\tPATH")
		if err != nil {
			return err
		}

		for _, r := range cfg.Rules {
			_, err = fmt.Fprintf(w, "%s\t%s\n", r.Profile, r.Path)
			if err != nil {
				return err
			}
		}
		err = w.Flush()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	ruleCmd.AddCommand(ruleListCmd)
}
