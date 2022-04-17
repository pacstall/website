package git

import (
	"os"
	"os/exec"
)

func hardResetAndPull(path string) error {
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

func clonePrograms(path, url string) error {
	cmd := exec.Command("git", "clone", url, path)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func RefreshPrograms(path, url string) error {
	if err := hardResetAndPull(path); err == nil {
		return nil
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	if err := clonePrograms(path, url); err != nil {
		return err
	}

	return nil
}
