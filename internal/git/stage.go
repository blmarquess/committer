package git

import (
	"os/exec"
	"strings"
)

func IsStageEmpty() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	outputString := string(output)
	return strings.TrimSpace(outputString) == "", nil
}

func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outputString := string(output)
	lines := strings.Split(strings.TrimSpace(outputString), "\n")
	return lines, nil
}
