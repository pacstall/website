package file

import (
	"io/fs"
	"log"
	"os"
)

var CreateTempDirectory = createTempDirectory

func createTempDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, fs.FileMode(int(0777))); err != nil {
			log.Printf("Failed to create temp dir '%v'\n%v", path, err)
			return err
		}

		log.Printf("Created fresh temp dir '%v'\n", path)
	} else {
		if err := os.RemoveAll(path); err != nil {
			log.Printf("Failed to remove existing temp dir '%v'\n", path)
			return err
		}

		log.Printf("Removed existing temp dir '%v'\n", path)
		return createTempDirectory(path)
	}

	return nil
}
