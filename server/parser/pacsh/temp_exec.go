package pacsh

import (
	"io/fs"
	"os"
	"os/exec"
	"path"

	"pacstall.dev/webserver/log"
)

var createFile = os.Create
var joinPaths = path.Join
var execCommand = exec.Command

var CreateTempExecutable = createTempExecutable

func createTempExecutable(dirPath, fileName string, content []byte) (string, error) {
	tmpFile, err := createFile(joinPaths(dirPath, fileName))

	if err != nil {
		log.Error.Printf("Failed to create temporary file '%v' in dir '%v'\n", fileName, dirPath)
		return "", err
	}
	defer tmpFile.Close()
	tmpPath := tmpFile.Name()

	defer func() {
		cmd := execCommand("chmod", "+rwx", fileName)
		cmd.Dir = dirPath
		if err := cmd.Run(); err != nil {
			log.Error.Printf("Failed to chmod temporary file '%v' in dir '%v'\n", fileName, dirPath)
		}
	}()

	if _, err = tmpFile.Write([]byte(content)); err != nil {
		log.Error.Printf("Failed to write to file '%v'\n%v", tmpPath, err)
		return "", err
	}

	if err := tmpFile.Chmod(fs.FileMode(int(0777))); err != nil {
		log.Error.Printf("Failed to chmod file '%v'\n%v", tmpPath, err)
		return "", err
	}

	return tmpPath, nil
}
