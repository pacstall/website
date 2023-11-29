package git

import (
	"os"
	"os/exec"

	"github.com/joomcode/errorx"
)

func hardResetAndPull(path string) error {
	cmd := exec.Command("git", "reset", "--hard", "HEAD")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	// https://stackoverflow.com/a/41081908/13449010
	cmd = exec.Command("git", "fetch")
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
	cmd := exec.Command("git", "clone", url, path)
	if err := cmd.Run(); err != nil {
		return errorx.Decorate(err, "failed to run git clone command")
	}

	return nil
}

func RefreshPrograms(path, url string) error {
	if err := hardResetAndPull(path); err == nil {
		return nil
	}

	if err := os.RemoveAll(path); err != nil {
		return errorx.Decorate(err, "failed to remove directory '%v'", path)
	}

	if err := clonePrograms(path, url); err != nil {
		return errorx.Decorate(err, "failed to clone repository '%v'", url)
	}

	return nil
}
