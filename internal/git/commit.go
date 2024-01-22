package git

import "os/exec"

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
