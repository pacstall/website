package internal_test

import (
	"testing"

	"pacstall.dev/webserver/internal/pacnexus/types/pac/parser/pacsh/internal"
)

func assertGitSourceInfoEquals(t *testing.T, expected, actual internal.GitSourceInfo) {
	t.Helper()

	if len(expected.Urls) != len(actual.Urls) {
		t.Errorf("expected url %+v but got %+v", expected.Urls, actual.Urls)
	}

	for idx, expectedUrl := range expected.Urls {
		if expectedUrl != actual.Urls[idx] {
			t.Errorf("expected url %+v but got %+v", expected.Urls, actual.Urls)
			return
		}
	}

	if expected.Branch != actual.Branch {
		t.Errorf("expected branch '%v' but got '%v'", expected.Branch, actual.Branch)
	}

	if expected.Tag != actual.Tag {
		t.Errorf("expected tag '%v' but got '%v'", expected.Tag, actual.Tag)
	}

	if expected.Commit != actual.Commit {
		t.Errorf("expected commit '%v' but got '%v'", expected.Commit, actual.Commit)
	}
}

func Test_ExtractGitSourceInformation_Simple(t *testing.T) {
	actual := internal.ExtractGitSourceInformation("git://git.deluge-torrent.org/deluge.git")
	expected := internal.GitSourceInfo{
		Urls: []string{"git://git.deluge-torrent.org/deluge.git"},
	}

	assertGitSourceInfoEquals(t, expected, actual)
}

func Test_ExtractGitSourceInformation_Tag(t *testing.T) {
	actual := internal.ExtractGitSourceInformation("git://git.deluge-torrent.org/deluge.git#tag=example_tag")
	expected := internal.GitSourceInfo{
		Urls: []string{"git://git.deluge-torrent.org/deluge.git"},
		Tag:  "example_tag",
	}

	assertGitSourceInfoEquals(t, expected, actual)
}

func Test_ExtractGitSourceInformation_Branch(t *testing.T) {
	actual := internal.ExtractGitSourceInformation("git://git.deluge-torrent.org/deluge.git#branch=develop")
	expected := internal.GitSourceInfo{
		Urls:   []string{"git://git.deluge-torrent.org/deluge.git"},
		Branch: "develop",
	}

	assertGitSourceInfoEquals(t, expected, actual)
}

func Test_ExtractGitSourceInformation_Commit(t *testing.T) {
	actual := internal.ExtractGitSourceInformation("git://git.deluge-torrent.org/deluge.git#commit=a0b1b2d3e4")
	expected := internal.GitSourceInfo{
		Urls:   []string{"git://git.deluge-torrent.org/deluge.git"},
		Commit: "a0b1b2d3e4",
	}

	assertGitSourceInfoEquals(t, expected, actual)
}

func Test_ExtractGitSourceInformation_GitPlusHttps(t *testing.T) {
	actual := internal.ExtractGitSourceInformation("git+https://git.deluge-torrent.org/deluge.git#commit=a0b1b2d3e4")
	expected := internal.GitSourceInfo{
		Urls:   []string{"https://git.deluge-torrent.org/deluge.git", "git://git.deluge-torrent.org/deluge.git"},
		Commit: "a0b1b2d3e4",
	}

	assertGitSourceInfoEquals(t, expected, actual)
}

func Test_ExtractGitSourceInformation_WithNamePrefix(t *testing.T) {
	actual := internal.ExtractGitSourceInformation("someName::git+https://git.deluge-torrent.org/deluge.git#commit=a0b1b2d3e4")
	expected := internal.GitSourceInfo{
		Urls:   []string{"https://git.deluge-torrent.org/deluge.git", "git://git.deluge-torrent.org/deluge.git"},
		Commit: "a0b1b2d3e4",
	}

	assertGitSourceInfoEquals(t, expected, actual)
}

func Test_ParseGitPackageVersion(t *testing.T) {
	pkgver, err := internal.NewTestGitSources([]string{"git://git.deluge-torrent.org/deluge.git#commit=testing123"}).ParseGitPackageVersion()
	if err != nil {
		t.Error(err)
		return
	}

	if pkgver != "testing1" {
		t.Errorf("expected 'testing1' but got '%v'", pkgver)
	}
}

func Test_ParseGitPackageVersion_NoCommit(t *testing.T) {
	pkgver, err := internal.NewTestGitSources([]string{"git://git.deluge-torrent.org/deluge.git#branch=testing123"}).ParseGitPackageVersion()
	if err != nil {
		t.Error(err)
		return
	}

	if pkgver != "a0b1c2d3" {
		t.Errorf("expected 'a0b1c2d3' but got '%v'", pkgver)
	}
}
