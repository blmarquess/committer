package git

import (
	"os/exec"
	"strings"
)

func GitPR() (bool, error) {
	cmd:= exec.Command("git", "symbolic-ref", "--short" "HEAD")
	branchOutput, err := cmd.Output()
	if err != nil {
		println(string(err))
		return false, err
	}
	actualBranchName := strings.TrimSpace(string(branchOutput))

	gitPRCommand := exec.Command("git", "push", "--set-upstream", "origin", actualBranchName)
	output, err := gitPRCommand.Output()
	if err != nil {
		println(string(err))
		return false, err
	}
	println(string(output))
	return true, nil
}