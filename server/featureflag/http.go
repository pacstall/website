package featureflag

import (
	"net/http"

	"pacstall.dev/website/config"
	"pacstall.dev/website/listener"
)

func GetFeatureFlags(w http.ResponseWriter, req *http.Request) {
	listener.ApplyHeaders("default", w, req)

	listener.Json(w, struct {
		PackageListPageDisabled    bool `json:"packageListPageDisabled"`
		PackageDetailsPageDisabled bool `json:"packageDetailsPageDisabled"`
		OldSyntax                  bool `json:"oldSyntax"`
	}{
		PackageListPageDisabled:    config.Config.FeatureFlags.PackageListPageDisabled,
		PackageDetailsPageDisabled: config.Config.FeatureFlags.PackageDetailsPageDisabled,
		OldSyntax:                  config.Config.FeatureFlags.OldSyntax,
	})
}
