package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetValue(key string) string {
	cmd := exec.Command("git", "config", "--get", key)
	var out bytes.Buffer
	cmd.Stdout = &out

	_ = cmd.Run()

	return strings.TrimSpace(out.String())
}
