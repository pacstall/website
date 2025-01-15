package pacsh

import (
	"github.com/pacstall/go-srcinfo"
	"pacstall.dev/webserver/types/pac"
)

func ParsePacOutput(data []byte) ([]*pac.Script, error) {
	var scripts []*pac.Script
	out, err := srcinfo.Parse(string(data))
	if err != nil {
		return nil, err
	}

	scripts = pac.FromSrcInfo(*out)
	for idx := range scripts {
        scripts[idx].PrettyName = getPrettyName(scripts[idx])
    }

	return scripts, nil
}
