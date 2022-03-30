package file

import (
	"os"
	"strings"
	"testing"
	"time"

	"pacstall.dev/webserver/types/list"
)

func cleanup() {
	statFile = os.Stat
	removeAll = os.RemoveAll
	makeDir = os.Mkdir
	removeFile = os.Remove
}

type testFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (t testFileInfo) Name() string {
	return t.name
}

func (t testFileInfo) Size() int64 {
	return t.size
}

func (t testFileInfo) Mode() os.FileMode {
	return t.mode
}

func (t testFileInfo) ModTime() time.Time {
	return t.modTime
}

func (t testFileInfo) IsDir() bool {
	return t.isDir
}

func (t testFileInfo) Sys() interface{} {
	return nil
}

func Test_CreateTempDirectory_NoExisting(t *testing.T) {
	defer cleanup()

	makeDirCalled := 0
	removeDirCalled := 0
	statFileCalled := 0

	statFile = func(path string) (os.FileInfo, error) {
		if statFileCalled == 1 {
			statFileCalled += 1
			return nil, os.ErrNotExist
		}

		statFileCalled += 1
		name, _ := list.From(strings.Split(path, "/")).Last()
		return testFileInfo{
			name:    name,
			size:    0,
			mode:    0777,
			modTime: time.Now(),
			isDir:   true,
		}, nil
	}

	makeDir = func(path string, perm os.FileMode) error {
		makeDirCalled += 1
		return nil
	}

	removeAll = func(path string) error {
		removeDirCalled += 1
		return nil
	}

	err := CreateTempDirectory("/tmp")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if removeDirCalled != 1 {
		t.Error("Expected removeAll to be called 1 time but was called", removeDirCalled)
	}

	if makeDirCalled != 1 {
		t.Error("Expected makeDir to be called 1 time but was called", makeDirCalled)
	}

	if statFileCalled != 2 {
		t.Error("Expected statFile to be called 2 times but was called", statFileCalled)
	}
}

func Test_CreateTempDirectory_AlreadyExisting(t *testing.T) {
	defer cleanup()

	makeDirCalled := 0
	removeDirCalled := 0
	statFileCalled := 0

	statFile = func(path string) (os.FileInfo, error) {
		if statFileCalled == 0 {
			statFileCalled += 1
			return nil, os.ErrNotExist
		}

		statFileCalled += 1
		name, _ := list.From(strings.Split(path, "/")).Last()
		return testFileInfo{
			name:    name,
			size:    0,
			mode:    0777,
			modTime: time.Now(),
			isDir:   true,
		}, nil
	}

	makeDir = func(path string, perm os.FileMode) error {
		makeDirCalled += 1
		return nil
	}

	removeAll = func(path string) error {
		removeDirCalled += 1
		return nil
	}

	err := CreateTempDirectory("/tmp")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if removeDirCalled != 0 {
		t.Error("Expected removeAll to be called 0 times but was called", removeDirCalled)
	}

	if makeDirCalled != 1 {
		t.Error("Expected makeDir to be called 1 time but was called", makeDirCalled)
	}

	if statFileCalled != 1 {
		t.Error("Expected statFile to be called 2 times but was called", statFileCalled)
	}
}
