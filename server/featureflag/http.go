package featureflag

import (
	"net/http"

	"pacstall.dev/website/config"
	"pacstall.dev/website/listener"
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
		OldSyntax: config.Config.FeatureFlags.OldSyntax,
		PackageDetailsPage: packageDetailsPageFeatureFlagsJson{
			LastUpdated:     config.Config.FeatureFlags.PackageDetailsPage.LastUpdated,
			Votes:           config.Config.FeatureFlags.PackageDetailsPage.Votes,
			Popularity:      config.Config.FeatureFlags.PackageDetailsPage.Popularity,
			InstallProtocol: config.Config.FeatureFlags.PackageDetailsPage.InstallProtocol,
			Comments:        config.Config.FeatureFlags.PackageDetailsPage.Comments,
		},
	}

	listener.Json(w, response)
}
