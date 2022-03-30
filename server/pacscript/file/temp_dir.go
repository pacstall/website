package file

import (
	"io/fs"
	"log"
	"os"
)

var CreateTempDirectory = createTempDirectory

var statFile = os.Stat
var removeAll = os.RemoveAll
var makeDir = os.Mkdir

func createTempDirectory(path string) error {
	if _, err := statFile(path); os.IsNotExist(err) {
		if err = makeDir(path, fs.FileMode(int(0777))); err != nil {
			log.Printf("Failed to create temp dir '%v'\n%v", path, err)
			return err
		}

	} else {
		if err := removeAll(path); err != nil {
			log.Printf("Failed to remove existing temp dir '%v'\n", path)
			return err
		}

		return createTempDirectory(path)
	}

	return nil
}
