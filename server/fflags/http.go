package fflags

import (
	"net/http"

	"pacstall.dev/website/cfg"
	"pacstall.dev/website/svlib"
)

func GetFeatureFlags(w http.ResponseWriter, req *http.Request) {
	svlib.ApplyHeaders("default", w, req)

	svlib.Json(w, struct {
		PackageListPageDisabled    bool `json:"packageListPageDisabled"`
		PackageDetailsPageDisabled bool `json:"packageDetailsPageDisabled"`
		OldSyntax                  bool `json:"oldSyntax"`
	}{
		PackageListPageDisabled:    cfg.Config.FeatureFlags.PackageListPageDisabled,
		PackageDetailsPageDisabled: cfg.Config.FeatureFlags.PackageDetailsPageDisabled,
		OldSyntax:                  cfg.Config.FeatureFlags.OldSyntax,
	})
}
