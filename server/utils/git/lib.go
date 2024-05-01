package git

import (
	"os"
	"os/exec"
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/log"
)

func hardResetAndPull(path, branch string) error {
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

	if err := checkoutBranch(path, branch); err != nil {
		return err
	}

	cmd = exec.Command("git", "pull")
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func clonePrograms(path, url, branch string) error {
	cmd := exec.Command("git", "clone", url, path)
	if err := cmd.Run(); err != nil {
		return errorx.Decorate(err, "failed to run git clone command")
	}

	if err := checkoutBranch(path, branch); err != nil {
		return err
	}

	return nil
}

func getCurrentBranch(path string) (string, error) {
	cmd := exec.Command("git", "name-rev", "--name-only", "HEAD")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return "", nil
	}

	return strings.TrimSpace(string(out)), nil
}

func checkoutBranch(path, branch string) error {
	currentBranch, err := getCurrentBranch(path)
	if err != nil {
		return errorx.Decorate(err, "failed to read current branch name")
	}

	if currentBranch != branch {
		log.Warn("programs repository is on the wrong branch '%v'. checking out branch '%v'", currentBranch, branch)
	} else {
		log.Info("programs repository is using branch: %v", currentBranch)
		return nil
	}

	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return errorx.Decorate(err, "failed to checkout git branch '%v'", branch)
	}

	log.Info("got branch checkout output: %v", string(out))

	return nil
}

// Returns the remote Git (short) commit hash
func GetRemoteCommitHash(url, branchOrTag string) (string, error) {
	cmd := exec.Command("git", "ls-remote", url, branchOrTag)

	if bytes, err := cmd.Output(); err != nil {
		return "", errorx.ExternalError.Wrap(err, "failed to fetch git commit hash from source '%v' branch/tag '%v'", url, branchOrTag)
	} else if len(string(bytes)) < 8 {
		return "", errorx.ExternalError.New("commit hash '%v' has less than 8 characters. source '%v' branch/tag '%v'", string(bytes), url, branchOrTag)
	} else {
		return strings.TrimSpace(string(bytes))[0:8], nil
	}
}

func RefreshPrograms(path, url, branch string) error {
	if err := hardResetAndPull(path, branch); err == nil {
		return nil
	}

	if err := os.RemoveAll(path); err != nil {
		return errorx.Decorate(err, "failed to remove directory '%v'", path)
	}

	if err := clonePrograms(path, url, branch); err != nil {
		return errorx.Decorate(err, "failed to clone repository '%v'", url)
	}

	return nil
}
