package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/spf13/cobra"
)

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "get list profle",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		_, err = fmt.Fprintln(w, "PROFILE\tNAME\tEMAIL")
		if err != nil {
			return err
		}

		for id, p := range cfg.Profiles {
			_, err = fmt.Fprintf(w, "%s\t%s\t%s\n", id, p.Name, p.Email)
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
	profileCmd.AddCommand(profileListCmd)
}
