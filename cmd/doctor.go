package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sen9kuni/hats/internal/config"
	"github.com/sen9kuni/hats/internal/engine"
	"github.com/sen9kuni/hats/internal/utils"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check for misconfiguration and system healt",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Running Hats diagnotics...")

		hasError := false

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not get home directory: %w", err)
		}
		hatsDir := filepath.Join(home, ".config", "hats")

		// NOTE: Check Hats directory
		if _, err := os.Stat(hatsDir); os.IsNotExist(err) {
			fmt.Printf("[Directory] Base directory missing: %s\n", hatsDir)
			fmt.Println("\tFix: Run 'hats init' to boostrap the environment.")
			return nil
		}
		fmt.Printf("[Directory] Found base directory: %s\n", hatsDir)

		// NOTE: Check state file hats.toml
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("[State] Failed to parse hats.toml: %v\n", err)
			hasError = true
		} else {
			fmt.Println("[State] hats.toml parsed successfully.")
		}

		// NOTE: Check global git integeration
		globalConfigPath := filepath.Join(home, ".gitconfig")
		content, err := os.ReadFile(globalConfigPath)
		if err != nil {
			fmt.Printf("[Git Hook] Could not read global ~/.gitconfig\n")
			hasError = true
		} else {
			expectedInclude := filepath.Join(hatsDir, config.IncludesFileName)
			if strings.Contains(string(content), engine.FormatToTilde(expectedInclude)) {
				fmt.Println("[Git Hook] Global includeIf hook is correctly installed")
			} else {
				fmt.Printf("[Git Hook] Missing include hook in ~/.gitconfig\n")
				fmt.Println("	Fix: Run 'hats init' to reinstall the global hook.")
				hasError = true
			}
		}

		// NOTE: Check profile synchronization
		if cfg != nil {
			for id := range cfg.Profiles {
				profilePath := filepath.Join(hatsDir, config.ProfileDirName, id+utils.GitFileName)
				if _, err := os.Stat(profilePath); os.IsNotExist(err) {
					fmt.Printf("[Profiles] Mising generated file for profile: '%s'\n", id)
					hasError = true
				}
			}

			if !hasError && len(cfg.Profiles) > 0 {
				fmt.Printf("[Profiles]: All %d profiles are properly synchronized.\n", len(cfg.Profiles))
			} else if len(cfg.Profiles) == 0 {
				fmt.Println("[Profiles] No profiles configured yet.")
			}
		}

		fmt.Println("\n--------------------------------------------------")
		if hasError {
			fmt.Println("Issue found. See the fix instructions above.")
		} else {
			fmt.Println("System is 100% healthy! Hats is ready to work.")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(doctorCmd)
}
