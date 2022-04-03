package config

type FeatureFlagsConfig struct {
	OldSyntax          bool
	PackageDetailsPage PackageDetailsPageFeatureFlags
}

type PackageDetailsPageFeatureFlags struct {
	LastUpdated     bool
	Votes           bool
	Popularity      bool
	InstallProtocol bool
	Comments        bool
}

var FeatureFlags = FeatureFlagsConfig{}

func setFeatureFlags(conf tomlFeatureFlagsConfig) {
	FeatureFlags.OldSyntax = conf.OldSyntax
	FeatureFlags.PackageDetailsPage = PackageDetailsPageFeatureFlags{
		LastUpdated:     conf.PackageDetailsPage.LastUpdated,
		Votes:           conf.PackageDetailsPage.Votes,
		Popularity:      conf.PackageDetailsPage.Popularity,
		InstallProtocol: conf.PackageDetailsPage.InstallProtocol,
		Comments:        conf.PackageDetailsPage.Comments,
	}
}
