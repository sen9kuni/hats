package cmd

import (
	"fmt"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var ruleRemoveCmd = &cobra.Command{
	Use:   "remove [path-apply-rule]",
	Short: "remove a rule",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rawPath := args[0]
		targetPath := engine.FormatToTilde(rawPath)

		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}
		updatedRules := []config.Rule{}
		removed := false

		for _, r := range cfg.Rules {
			if r.Path == targetPath {
				removed = true
				continue
			}
			updatedRules = append(updatedRules, r)
		}

		if !removed {
			return fmt.Errorf("no rule found for path: %s", targetPath)
		}

		cfg.Rules = updatedRules

		err = config.Save(cfg)
		if err != nil {
			return err
		}
		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("rule deleted, but failed to apply to git: %w", err)
		}
		fmt.Printf("removed rule for: %s\n", targetPath)
		return nil
	},
}

func init() {
	ruleCmd.AddCommand(ruleRemoveCmd)
}
