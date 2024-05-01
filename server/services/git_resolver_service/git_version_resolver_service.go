package grs

import (
	"strings"

	"github.com/joomcode/errorx"
)

type GitVersionResolverImpl struct {
	commitResolver GitCommitResolver
}

func New(commitResolver GitCommitResolver) *GitVersionResolverImpl {
	r := &GitVersionResolverImpl{}

	r.commitResolver = commitResolver

	return r
}

func (s *GitVersionResolverImpl) ParseGitPackageVersion(sources []string) (string, error) {
	if len([]string(sources)) == 0 {
		// Nothing to parse
		return "", nil
	}

	primarySource := []string(sources)[0]
	// Only keep the url part of "mycoolname::git+https://github.com/me/project.git#branch=coolfeature"
	sourceInfo := extractGitSourceInformation(primarySource)

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
			out, err := s.tryGetCommitHashFromAnySource(sourceInfo.Urls, sourceInfo.Tag)
			return out, err
		}
	} else if sourceInfo.Branch != "" {
		calculateCommit = func() (string, error) {
			out, err := s.tryGetCommitHashFromAnySource(sourceInfo.Urls, sourceInfo.Branch)
			return out, err
		}
	} else {
		calculateCommit = func() (string, error) {
			out, err := s.tryGetCommitHashFromAnySource(sourceInfo.Urls, "HEAD")
			return out, err
		}
	}

	commitHash, err := calculateCommit()
	return commitHash, err
}

func (s *GitVersionResolverImpl) tryGetCommitHashFromAnySource(urls []string, ref string) (string, error) {
	errors := []error{}

	for _, url := range urls {
		out, err := s.commitResolver.GetCommitHash(url, ref)
		if err == nil {
			return out, nil
		}

		errors = append(errors, err)
	}

	return "", errorx.DecorateMany("failed to get commit hash from any source", errors...)
}

type gitSourceInfo struct {
	Urls   []string
	Branch string
	Tag    string
	Commit string
}

func extractGitSourceInformation(source string) gitSourceInfo {
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

	urls := []string{}
	if strings.Contains(sourceUrl, "git+https://") {
		httpsUrl := strings.ReplaceAll(sourceUrl, "git+https://", "https://")
		gitUrl := strings.ReplaceAll(sourceUrl, "git+https://", "git://")
		urls = append(urls, httpsUrl, gitUrl)
	} else {
		urls = append(urls, sourceUrl)
	}

	return gitSourceInfo{
		Urls:   urls,
		Tag:    tag,
		Branch: branch,
		Commit: commit,
	}
}
