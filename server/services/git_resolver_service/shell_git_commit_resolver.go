package grs

import (
	"fmt"
	"time"

	"pacstall.dev/webserver/utils/git"
	"pacstall.dev/webserver/utils/parallelism/timeout"
)

type ShellGitCommitResolver struct{}

func NewShellGitCommitResolver() *ShellGitCommitResolver {
	r := &ShellGitCommitResolver{}

	return r
}

func (*ShellGitCommitResolver) GetCommitHash(url, ref string) (string, error) {
	version, err := timeout.Run(fmt.Sprintf("commit-hash/%v/%v", url, ref), func() (string, error) {
		version, err := git.GetRemoteCommitHash(url, ref)
		return version, err
	}, 3*time.Second)
	return version, err
}
