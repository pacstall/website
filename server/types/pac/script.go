package pac

import "time"

type updateStatus struct {
	Unknown UpdateStatusValue
	Latest  UpdateStatusValue
	Minor   UpdateStatusValue
	Major   UpdateStatusValue
	Patch   UpdateStatusValue
}

var UpdateStatus = updateStatus{
	Major:   3,
	Minor:   2,
	Patch:   1,
	Latest:  0,
	Unknown: -1,
}

type UpdateStatusValue = int

type Script struct {
	Name                 string    `json:"name"`
	PrettyName           string    `json:"prettyName"`
	Version              string    `json:"version"`
	LatestVersion        *string   `json:"latestVersion"`
	PackageName          string    `json:"packageName"`
	Maintainer           string    `json:"maintainer"`
	Description          string    `json:"description"`
	URL                  string    `json:"url"`
	RuntimeDependencies  []string  `json:"runtimeDependencies"`
	BuildDependencies    []string  `json:"buildDependencies"`
	OptionalDependencies []string  `json:"optionalDependencies"`
	Breaks               []string  `json:"breaks"`
	Gives                string    `json:"gives"`
	Replace              []string  `json:"replace"`
	Hash                 *string   `json:"hash"`
	PPA                  []string  `json:"ppa"`
	PacstallDependencies []string  `json:"pacstallDependencies"`
	Patch                []string  `json:"patch"`
	Repology             []string  `json:"repology"`
	RequiredBy           []string  `json:"requiredBy"`
	LastUpdatedAt        time.Time `json:"lastUpdatedAt"`
	UpdateStatus         int       `json:"updateStatus"` // enum UpdateStatus
}
