package pacsh

import (
	"os"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/pkg/common/log"
)

var removeFile = os.Remove
var ExecBash = execBash

func execBash(cwd, filename string, content []byte) (stdout []byte, err error) {
	tmpPath, err := CreateTempExecutable(cwd, filename, content)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to create temp executable")
	}
	defer removeFile(tmpPath)

	command := execCommand("bash", tmpPath)
	command.Env = append(command.Env, "CARCH=amd64")
	stdout, err = command.Output()
	if err != nil {
		bytes, _ := os.ReadFile(tmpPath)
		log.Debug("Failed to execute '%v'. %v\n%v", tmpPath, err, string(bytes))
		return nil, errorx.Decorate(err, "failed to execute '%v'", tmpPath)
	}

	return stdout, nil
}
