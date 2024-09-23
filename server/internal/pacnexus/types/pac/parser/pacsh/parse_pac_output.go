package pacsh

import (
	"github.com/pacstall/go-srcinfo"
	"pacstall.dev/webserver/internal/pacnexus/types/pac"
)

func ParsePacOutput(data []byte) (*pac.Script, error) {
	out, err := srcinfo.Parse(string(data))
	if err != nil {
		return nil, err
	}

	ps := pac.FromSrcInfo(*out)

	ps.PrettyName = getPrettyName(ps)

	return ps, nil
}
