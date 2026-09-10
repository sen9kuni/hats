package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var remoteListCmd = &cobra.Command{
	Use:   "list",
	Short: "get list remote",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		_, err = fmt.Fprintln(w, "PROFILE\tREMOTE URL")
		if err != nil {
			return err
		}

		for _, r := range cfg.RemoteRules {
			_, err = fmt.Fprintf(w, "%s\t%s\n", r.Profile, r.URL)
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
	remoteCmd.AddCommand(remoteListCmd)
}
