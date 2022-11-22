package pacsh

import (
	"io/fs"
	"os"

	"pacstall.dev/webserver/log"
)

var CreateTempDirectory = createTempDirectory

var statFile = os.Stat
var removeAll = os.RemoveAll
var makeDir = os.Mkdir

func createTempDirectory(path string) error {
	if _, err := statFile(path); os.IsNotExist(err) {
		if err = makeDir(path, fs.FileMode(int(0777))); err != nil {
			log.Error("Failed to create temp dir '%v'\n%v", path, err)
			return err
		}

	} else {
		if err := removeAll(path); err != nil {
			log.Error("Failed to remove existing temp dir '%v'", path)
			return err
		}

		return createTempDirectory(path)
	}

	return nil
}
