package cache

import "pacstall.dev/website/types"

var cache map[string][]*types.PackageInfo = make(map[string][]*types.PackageInfo)

func Get(etag string) ([]*types.PackageInfo, bool) {
	found, ok := cache[etag]
	return found, ok
}

func Set(etag string, pkgs []*types.PackageInfo) {
	cache[etag] = pkgs
}

func Invalidate() {
	cache = make(map[string][]*types.PackageInfo)
}
