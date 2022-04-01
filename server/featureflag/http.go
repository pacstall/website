package featureflag

import (
	"net/http"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/listener"
)

type packageDetailsPageFeatureFlagsJson struct {
	LastUpdated     bool `json:"lastUpdated"`
	Votes           bool `json:"votes"`
	Popularity      bool `json:"popularity"`
	InstallProtocol bool `json:"installProtocol"`
	Comments        bool `json:"comments"`
}

type featureFlagsJson struct {
	OldSyntax          bool                               `json:"oldSyntax"`
	PackageDetailsPage packageDetailsPageFeatureFlagsJson `json:"packageDetailsPage"`
}

func GetFeatureFlags(w http.ResponseWriter, req *http.Request) {
	listener.ApplyHeaders("default", w, req)

	response := featureFlagsJson{
		OldSyntax: config.FeatureFlags.OldSyntax,
		PackageDetailsPage: packageDetailsPageFeatureFlagsJson{
			LastUpdated:     config.FeatureFlags.PackageDetailsPage.LastUpdated,
			Votes:           config.FeatureFlags.PackageDetailsPage.Votes,
			Popularity:      config.FeatureFlags.PackageDetailsPage.Popularity,
			InstallProtocol: config.FeatureFlags.PackageDetailsPage.InstallProtocol,
			Comments:        config.FeatureFlags.PackageDetailsPage.Comments,
		},
	}

	listener.Json(w, response)
}
