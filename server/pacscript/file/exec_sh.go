package file

import (
	"log"
	"os"
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
		log.Printf("Failed to execute '%v'. %v\n%v", tmpPath, err, string(bytes))
		return
	}

	return stdout, nil
}
