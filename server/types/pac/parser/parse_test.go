package parser_test

import (
	"os"
	"path"
	"testing"

	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/parser"
	"pacstall.dev/webserver/types/pac/parser/pacsh"
)

var FIXTURES_DIR = func() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path.Join(dir, "../../../fixtures")
}()

var TEST_PROGRAMS_DIR = path.Join(FIXTURES_DIR, "test-programs")

func assertEquals(t *testing.T, what string, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("expected %v '%v', got '%v'", what, expected, actual)
	}
}

func assertArrayEquals(t *testing.T, what string, expected []string, actual []string) {
	if len(actual) != len(expected) {
		t.Errorf("expected %v '%v', got '%v'", what, expected, actual)
	}

	for idx := range expected {
		if expected[idx] != actual[idx] {
			t.Errorf("expected %v '%v', got '%v'", what, expected, actual)
		}
	}
}

func assertPacscriptEquals(t *testing.T, expected pac.Script, actual pac.Script) {
	assertEquals(t, "name", expected.Name, actual.Name)
	assertEquals(t, "package name", expected.PackageName, actual.PackageName)
	assertEquals(t, "maintainer", expected.Maintainer, actual.Maintainer)
	assertEquals(t, "description", expected.Description, actual.Description)
	assertEquals(t, "gives", expected.Gives, actual.Gives)
	assertEquals(t, "hash", *expected.Hash, *actual.Hash)
	assertEquals(t, "version", expected.Version, actual.Version)
	assertArrayEquals(t, "breaks", expected.Breaks, actual.Breaks)
	assertArrayEquals(t, "replace", expected.Replace, actual.Replace)
	assertEquals(t, "pretty name", expected.PrettyName, actual.PrettyName)
	assertEquals(t, "url", expected.URL, actual.URL)
	assertArrayEquals(t, "runtime dependencies", expected.RuntimeDependencies, actual.RuntimeDependencies)
	assertArrayEquals(t, "build dependencies", expected.BuildDependencies, actual.BuildDependencies)
	assertArrayEquals(t, "optional dependencies", expected.OptionalDependencies, actual.OptionalDependencies)
	assertArrayEquals(t, "pacstall dependencies", expected.PacstallDependencies, actual.PacstallDependencies)
	assertArrayEquals(t, "ppa", expected.PPA, actual.PPA)
	assertArrayEquals(t, "patch", expected.Patch, actual.Patch)
	assertArrayEquals(t, "required by", expected.RequiredBy, actual.RequiredBy)
	assertArrayEquals(t, "repology", expected.Repology, actual.Repology)
	assertEquals(t, "update status", expected.UpdateStatus, actual.UpdateStatus)
}

func Test_ParsePacscriptFile_Valid(t *testing.T) {
	if pacsh.CreateTempDirectory("./tmp") != nil {
		t.Errorf("failed to create temp directory")
		return
	}

	actual, err := parser.ParsePacscriptFile(TEST_PROGRAMS_DIR, "sample-valid-deb")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	hash := "10101010"
	expected := pac.Script{
		Name:                 "sample-valid-deb",
		PackageName:          "sample-valid",
		Maintainer:           "pacstall <test@pacstall.dev>",
		Description:          "Sample description",
		Gives:                "sample-valid",
		Hash:                 &hash,
		Version:              "1.0.0",
		Breaks:               []string{"breaks1", "breaks2"},
		Replace:              []string{"replaces1", "replaces2"},
		PrettyName:           "Sample Valid",
		URL:                  "https://example.com",
		RuntimeDependencies:  []string{"dep1", "dep2"},
		BuildDependencies:    []string{"go", "gcc"},
		OptionalDependencies: []string{"opt1", "opt2"},
		PacstallDependencies: []string{"pacdep1", "pacdep2"},
		PPA:                  []string{"ppa1", "ppa2"},
		Patch:                []string{"patch1", "patch2"},
		RequiredBy:           []string{},
		Repology:             []string{"project: sample-valid"},
		UpdateStatus:         pac.UpdateStatus.Unknown,
	}

	assertPacscriptEquals(t, expected, actual)
}

func Test_ParsePacscriptFile_WithPkgverFunc_Valid(t *testing.T) {
	if pacsh.CreateTempDirectory("./tmp") != nil {
		t.Errorf("failed to create temp directory")
		return
	}

	actual, err := parser.ParsePacscriptFile(TEST_PROGRAMS_DIR, "sample-valid-with-pkgver-func-deb")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	hash := "10101010"
	expected := pac.Script{
		Name:                 "sample-valid-deb",
		PackageName:          "sample-valid",
		Maintainer:           "pacstall <test@pacstall.dev>",
		Description:          "Sample description",
		Gives:                "sample-valid",
		Hash:                 &hash,
		Version:              "1.2.3",
		Breaks:               []string{"breaks1", "breaks2"},
		Replace:              []string{"replaces1", "replaces2"},
		PrettyName:           "Sample Valid",
		URL:                  "https://example.com",
		RuntimeDependencies:  []string{"dep1", "dep2"},
		BuildDependencies:    []string{"go", "gcc"},
		OptionalDependencies: []string{"opt1", "opt2"},
		PacstallDependencies: []string{"pacdep1", "pacdep2"},
		PPA:                  []string{"ppa1", "ppa2"},
		Patch:                []string{"patch1", "patch2"},
		RequiredBy:           []string{},
		Repology:             []string{"project: sample-valid"},
		UpdateStatus:         pac.UpdateStatus.Unknown,
	}

	assertPacscriptEquals(t, expected, actual)
}
