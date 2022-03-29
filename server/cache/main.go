package cache

import "pacstall.dev/webserver/types"

var cache map[string][]*types.Pacscript = make(map[string][]*types.Pacscript)

func Get(etag string) ([]*types.Pacscript, bool) {
	found, ok := cache[etag]
	return found, ok
}

func Set(etag string, pkgs []*types.Pacscript) {
	cache[etag] = pkgs
}

func Invalidate() {
	cache = make(map[string][]*types.Pacscript)
}
