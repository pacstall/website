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

	// https://stackoverflow.com/a/41081908/13449010
	cmd = exec.Command("git", "fetch", "--depth=1")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "reset", "--hard", "origin/HEAD")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "clean", "-dfx")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func clonePrograms(path, url string) error {
	cmd := exec.Command("git", "clone", "--depth=1", url, path)
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
