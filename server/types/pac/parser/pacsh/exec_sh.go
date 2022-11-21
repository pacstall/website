package pacsh

import (
	"os"

	"pacstall.dev/webserver/log"
)

var removeFile = os.Remove
var ExecBash = execBash

func execBash(cwd, filename string, pacscript []byte) (stdout []byte, err error) {
	tmpPath, err := CreateTempExecutable(cwd, filename, pacscript)
	if err != nil {
		return
	}
	defer removeFile(tmpPath)

	stdout, err = execCommand("bash", tmpPath).Output()
	if err != nil {
		bytes, _ := os.ReadFile(tmpPath)
		log.Debug.Printf("Failed to execute '%v'. %v\n%v\n", tmpPath, err, string(bytes))
		return
	}

	return stdout, nil
}
