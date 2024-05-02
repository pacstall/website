package pacsh

import (
	"io/fs"
	"os"
	"os/exec"
	"path"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/log"
)

var createFile = os.Create
var joinPaths = path.Join
var execCommand = exec.Command

var CreateTempExecutable = createTempExecutable

func createTempExecutable(dirPath, fileName string, content string) (string, error) {
	tmpFile, err := createFile(joinPaths(dirPath, fileName))

	if err != nil {
		return "", errorx.Decorate(err, "failed to create temporary file '%v' in dir '%v'", fileName, dirPath)
	}
	defer tmpFile.Close()
	tmpPath := tmpFile.Name()

	defer func() {
		cmd := execCommand("chmod", "+rwx", fileName)
		cmd.Dir = dirPath
		if err := cmd.Run(); err != nil {
			log.Error("%+v", errorx.Decorate(err, "failed to chmod temporary file '%v' in dir '%v'", fileName, dirPath))
		}
	}()

	if _, err = tmpFile.Write([]byte(content)); err != nil {
		return "", errorx.Decorate(err, "failed to write to file '%v'", tmpPath)
	}

	if err := tmpFile.Chmod(fs.FileMode(int(0777))); err != nil {
		return "", errorx.Decorate(err, "failed to chmod file '%v'", tmpPath)
	}

	return tmpPath, nil
}
