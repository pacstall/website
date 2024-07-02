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
	PrettyName           string    `json:"prettyName"`
	Version              string    `json:"version"`
	LatestVersion        *string   `json:"latestVersion"`
	PackageName          string    `json:"packageName"`
	Maintainers          []string  `json:"maintainers"`
	Description          string    `json:"description"`
	Source               []string  `json:"source"`
	RuntimeDependencies  []string  `json:"runtimeDependencies"`
	BuildDependencies    []string  `json:"buildDependencies"`
	OptionalDependencies []string  `json:"optionalDependencies"`
	Conflicts            []string  `json:"conflicts"`
	Breaks               []string  `json:"breaks"`
	Gives                string    `json:"gives"`
	Replaces             []string  `json:"replaces"`
	Hash                 *string   `json:"hash"`
	PPA                  []string  `json:"ppa"`
	PacstallDependencies []string  `json:"pacstallDependencies"`
	Patch                []string  `json:"patch"`
	Repology             []string  `json:"repology"`
	RequiredBy           []string  `json:"requiredBy"`
	LastUpdatedAt        time.Time `json:"lastUpdatedAt"`
	UpdateStatus         int       `json:"updateStatus"` // enum UpdateStatus
}
