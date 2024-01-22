package git

import (
	"os/exec"
	"strings"
)

func GetUntrackedFiles() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--others", "--exclude-standard")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outputString := string(output)
	lines := strings.Split(strings.TrimSpace(outputString), "\n")
	return lines, nil
}
