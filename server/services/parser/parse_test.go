package parser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"pacstall.dev/webserver/config"
	grs "pacstall.dev/webserver/services/git_resolver_service"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/services/parser"
	"pacstall.dev/webserver/services/parser/pacsh"
	"pacstall.dev/webserver/services/repology"
	"pacstall.dev/webserver/types/pac"
)

var FIXTURES_DIR = func() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path.Join(dir, "../../fixtures")
}()

var TEST_PROGRAMS_DIR = path.Join(FIXTURES_DIR, "test-programs")

type MockCommitResolver struct{}

func (*MockCommitResolver) GetCommitHash(url, ref string) (string, error) {
	return "", nil
}

func assertEquals(t *testing.T, what string, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("pacscript.%v: expected '%#v', got '%#v'", what, expected, actual)
	}
}

func assertArrayEquals(t *testing.T, what string, expected []string, actual []string) {
	if len(actual) == len(expected) && len(actual) == 0 {
		return
	}

	if len(actual) != len(expected) {
		t.Errorf("pacscript.%v expected len '%v', got len '%v' (expected '%#v', got '%#v')", what, len(expected), len(actual), expected, actual)
		return
	}

	for idx := range expected {
		if expected[idx] != actual[idx] {
			t.Errorf("pacscript.%v[%v] expected '%#v', got '%#v'", what, idx, expected, actual)
		}
	}
}

func assertPacscriptEquals(t *testing.T, expected pac.Script, actual pac.Script) {
	assertEquals(t, "package name", expected.PackageName, actual.PackageName)
	assertArrayEquals(t, "maintainers", expected.Maintainers, actual.Maintainers)
	assertEquals(t, "description", expected.Description, actual.Description)
	assertEquals(t, "gives", expected.Gives, actual.Gives)
	if expected.Hash != nil && actual.Hash == nil {
		t.Errorf("expected hash '%v', got nil", *expected.Hash)
	} else if expected.Hash == nil && actual.Hash != nil {
		t.Errorf("expected hash nil, got %v", *actual.Hash)
	} else if expected.Hash != nil && actual.Hash != nil {
		assertEquals(t, "hash", *expected.Hash, *actual.Hash)
	}
	if !strings.Contains(expected.PackageName, "-git") {
		assertEquals(t, "version", expected.Version, actual.Version)
	}
	assertArrayEquals(t, "breaks", expected.Breaks, actual.Breaks)
	assertArrayEquals(t, "conflicts", expected.Conflicts, actual.Conflicts)
	assertArrayEquals(t, "replaces", expected.Replaces, actual.Replaces)
	assertEquals(t, "pretty name", expected.PrettyName, actual.PrettyName)
	assertArrayEquals(t, "sources", expected.Source, actual.Source)
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

func loadSnapshot(snapshotPath string) (*pac.Script, error) {
	bytes, err := os.ReadFile(snapshotPath)
	if err != nil {
		return nil, err
	}

	var out pac.Script
	if err := json.Unmarshal(bytes, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func assertPacscriptMatchesSnapshot(t *testing.T, pkgname string) {
	t.Helper()

	if pacsh.CreateTempDirectory("./tmp") != nil {
		t.Errorf("failed to create temp directory")
		return
	}

	parserService := parser.New(
		config.PacstallProgramsConfiguration{
			ClonePath: "./programs",
		},
		config.ServerConfiguration{
			TempDir:      "./tmp",
			MaxOpenFiles: 100,
			Production:   false,
		},
		config.RepologyConfiguration{
			Enabled: false,
		},
		&repology.RepologyService{},
		grs.New(&MockCommitResolver{}),
		pkgcache.New(),
	)

	actual, err := parserService.ParsePacscriptFile(TEST_PROGRAMS_DIR, pkgname)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	snapshotPath := path.Join(TEST_PROGRAMS_DIR, "packages", pkgname, fmt.Sprintf("%v.snapshot.json", pkgname))
	expected, err := loadSnapshot(snapshotPath)
	if err != nil {
		bytes, err := json.Marshal(actual)
		if err != nil {
			t.Errorf("failed to serialize snapshot. %v", err)
			return
		}

		if err := os.WriteFile(snapshotPath, bytes, 0644); err != nil {
			t.Errorf("failed to write snapshot. %v", err)
			return
		}

		t.Errorf("missing snapshot. a new one has been generated. rerun tests")
		return
	}

	assertPacscriptEquals(t, *expected, actual)
}

func Test_PacscriptSnapshots(t *testing.T) {
	dirEntries, err := os.ReadDir(path.Join(TEST_PROGRAMS_DIR, "packages"))
	if err != nil {
		t.Errorf("failed to read test packages. %v", err)
		return
	}

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			continue
		}

		t.Logf("==> Running snapshot test for file: %v", dirEntry.Name())
		assertPacscriptMatchesSnapshot(t, dirEntry.Name())
	}
}
