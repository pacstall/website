package internal

import (
	"fmt"
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/types/pac/parser/git"
	"pacstall.dev/webserver/types/pac/parser/parallelism/timeout"
)

type GitSourceInfo struct {
	Url    string
	Branch string
	Tag    string
	Commit string
}

type GitSources struct {
	sources       []string
	getCommitHash func(string, string) (string, error)
}

func NewGitSources(sources []string) GitSources {
	return GitSources{
		sources: sources,
		getCommitHash: func(url, ref string) (string, error) {
			version, err := timeout.Run(fmt.Sprintf("commit-hash/%v/%v", url, ref), func() (string, error) {
				version, err := git.GetRemoteCommitHash(url, ref)
				return version, err
			}, 1000)
			return version, err
		},
	}
}

func NewTestGitSources(sources []string) GitSources {
	return GitSources{
		sources: sources,
		getCommitHash: func(url, target string) (string, error) {
			return "a0b1c2d3", nil
		},
	}
}

func ExtractGitSourceInformation(source string) GitSourceInfo {
	sourceUrl := ""
	if parts := strings.Split(source, "::"); len(parts) > 1 {
		sourceUrl = parts[1]
	} else {
		sourceUrl = parts[0]
	}

	commit := ""
	tag := ""
	branch := ""
	if strings.Contains(sourceUrl, "#branch=") {
		parts := strings.Split(sourceUrl, "#branch=")
		sourceUrl = parts[0]
		branch = parts[1]
	} else if strings.Contains(sourceUrl, "#tag=") {
		parts := strings.Split(sourceUrl, "#tag=")
		sourceUrl = parts[0]
		tag = parts[1]
	} else if strings.Contains(sourceUrl, "#commit=") {
		parts := strings.Split(sourceUrl, "#commit=")
		sourceUrl = parts[0]
		commit = parts[1]
	}

	sourceUrl = strings.ReplaceAll(sourceUrl, "git+https://", "git://")

	return GitSourceInfo{
		Url:    sourceUrl,
		Tag:    tag,
		Branch: branch,
		Commit: commit,
	}
}

func (s GitSources) ParseGitPackageVersion() (string, error) {
	if len([]string(s.sources)) == 0 {
		// Nothing to parse
		return "", nil
	}

	primarySource := []string(s.sources)[0]
	if !strings.Contains(primarySource, "git://") && !strings.Contains(primarySource, "git+") {
		// Probably not a Git source. Note that `git+` might appear somewhere in the source
		// and that might lead to a false positive.
		// TODO: Find a better way to check if the source is a git one.
		return "", nil
	}

	// Only keep the url part of "mycoolname::git+https://github.com/me/project.git#branch=coolfeature"
	sourceInfo := ExtractGitSourceInformation(primarySource)

	var calculateCommit func() (string, error) = nil

	if sourceInfo.Commit != "" {
		calculateCommit = func() (string, error) {
			if len(sourceInfo.Commit) < 8 {
				return "", errorx.AssertionFailed.New("expected commit hash '%v' to have more than 8 characters", sourceInfo.Commit)
			}

			return sourceInfo.Commit[0:8], nil
		}
	} else if sourceInfo.Tag != "" { // the following if branches look similar and could be merged but let's keep them this way for now.
		calculateCommit = func() (string, error) {
			out, err := s.getCommitHash(sourceInfo.Url, sourceInfo.Tag)
			return out, err
		}
	} else if sourceInfo.Branch != "" {
		calculateCommit = func() (string, error) {
			out, err := s.getCommitHash(sourceInfo.Url, sourceInfo.Branch)
			return out, err
		}
	} else {
		calculateCommit = func() (string, error) {
			out, err := s.getCommitHash(sourceInfo.Url, "HEAD")
			return out, err
		}
	}

	commitHash, err := calculateCommit()
	return commitHash, err
}
