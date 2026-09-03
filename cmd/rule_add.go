package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/spf13/cobra"
)

var ruleAddCmd = &cobra.Command{
	Use:   "add [profile-name] [directory-path]",
	Short: "Apply a profile to a specific directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		profileID := args[0]
		rawPath := args[1]

		hatsDir, err := config.GetPath()
		if err != nil {
			return err
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if _, exists := cfg.Profiles[profileID]; !exists {
			return fmt.Errorf("profile '%s' does not exists", profileID)
		}

		cleanPath := filepath.Clean(rawPath)
		if !strings.HasSuffix(cleanPath, "/") {
			cleanPath += "/"
		}

		cleanPath = engine.FormatToTilde(cleanPath)

		ruleUpdated := false
		for i, rule := range cfg.Rules {
			if rule.Path == cleanPath {
				cfg.Rules[i].Profile = profileID
				ruleUpdated = true
				break
			}
		}

		if !ruleUpdated {
			cfg.Rules = append(cfg.Rules, config.Rule{
				Profile: profileID,
				Path:    cleanPath,
			})
		}

		if err := config.Save(cfg); err != nil {
			return err
		}

		if err := engine.Sync(cfg, hatsDir); err != nil {
			return fmt.Errorf("rule saved, but failed to apply to git: %w", err)
		}

		fmt.Printf("Rule added: %s -> %s\n", cleanPath, profileID)
		return nil
	},
}

func init() {
	ruleCmd.AddCommand(ruleAddCmd)
}
