// Package engine for Logic for `apply`: taking Config and generating Git files
package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sen9kuni/hats/internal/config"
)

func FormatToTilde(rawPath string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return rawPath // fallback if we can't find home
	}

	// If the path starts with /Users/yourname, replace that part with ~
	if strings.HasPrefix(rawPath, home) {
		return strings.Replace(rawPath, home, "~", 1)
	}

	return rawPath
}

func GenerateProfileConfig(p config.Profile) string {
	var sb strings.Builder

	// write the user section
	fmt.Fprintf(&sb, "[user]\n")
	fmt.Fprintf(&sb, "\tname = %s\n", p.Name)
	fmt.Fprintf(&sb, "\temail = %s\n", p.Email)

	// check if the signing have value
	if p.SigningKey != "" {
		fmt.Fprintf(&sb, "\tsigningKey = %s\n", p.SigningKey)
		fmt.Fprintf(&sb, "[commit]\n")
		fmt.Fprintf(&sb, "\tgpgsign = true\n")
	}
	return sb.String()
}

func GenerateIncludesConfig(rules []config.Rule, profileDir string) string {
	// sort the longest string path on last
	sort.Slice(rules, func(i, j int) bool {
		return len(rules[i].Path) < len(rules[j].Path)
	})

	var sb strings.Builder
	for _, rule := range rules {
		gitdir := rule.Path

		if !strings.HasSuffix(gitdir, "/") {
			gitdir += "/"
		}

		profileDir = FormatToTilde(profileDir)
		profileFilePath := filepath.Join(profileDir, rule.Profile+".gitconfig")

		fmt.Fprintf(&sb, "[includeIf \"gitdir/i:%s\"]\n", gitdir)
		fmt.Fprintf(&sb, "\tpath = %s\n\n", profileFilePath)
	}

	return sb.String()
}

func GenerateIncludesRemoteConfig(remotes []config.Remote, profileDir string) string {
	var sb strings.Builder

	for _, remote := range remotes {
		profileDir = FormatToTilde(profileDir)
		profileFilePath := filepath.Join(profileDir, remote.Profile+".gitconfig")

		fmt.Fprintf(&sb, "[includeIf \"hasconfig:remote.*.url:%s\"]\n", remote.URL)
		fmt.Fprintf(&sb, "\tpath = %s\n\n", profileFilePath)
	}
	return sb.String()
}
