package pacsh

import (
	"os"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/log"
)

var removeFile = os.Remove
var ExecBash = execBash

func execBash(cwd, filename string, content string) (string, error) {
	tmpPath, err := CreateTempExecutable(cwd, filename, content)
	if err != nil {
		return "", errorx.Decorate(err, "failed to create temp executable")
	}
	defer removeFile(tmpPath)

	command := execCommand("bash", tmpPath)
	command.Env = append(command.Env, "CARCH=amd64")
	stdoutBytes, err := command.Output()
	if err != nil {
		bytes, _ := os.ReadFile(tmpPath)
		log.Debug("Failed to execute '%v'. %v\n%v", tmpPath, err, string(bytes))
		return "", errorx.Decorate(err, "failed to execute '%v'", tmpPath)
	}

	return string(stdoutBytes), nil
}
