package cache

import "pacstall.dev/webserver/types/pac"

var cache map[string][]*pac.Script = make(map[string][]*pac.Script)

func Get(etag string) ([]*pac.Script, bool) {
	found, ok := cache[etag]
	return found, ok
}

func Set(etag string, pkgs []*pac.Script) {
	cache[etag] = pkgs
}

func Invalidate() {
	cache = make(map[string][]*pac.Script)
}
