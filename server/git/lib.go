package git

import "os/exec"

func HardResetAndPull(path string) error {
	cmd := exec.Command("git", "reset", "--hard", "HEAD")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "fetch")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "pull")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
