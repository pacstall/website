package pacsh

import (
	"io/fs"
	"os"

	"github.com/joomcode/errorx"
)

var CreateTempDirectory = createTempDirectory

var statFile = os.Stat
var removeAll = os.RemoveAll
var makeDir = os.Mkdir

func createTempDirectory(path string) error {
	if _, err := statFile(path); os.IsNotExist(err) {
		if err = makeDir(path, fs.FileMode(int(0777))); err != nil {
			return errorx.Decorate(err, "failed to create temp dir '%v'", path)
		}

	} else {
		if err := removeAll(path); err != nil {
			return errorx.Decorate(err, "failed to remove existing temp dir '%v'", path)
		}

		return createTempDirectory(path)
	}

	return nil
}
