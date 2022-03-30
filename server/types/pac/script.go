package pac

type updateStatus struct {
	Unknown int
	Latest  int
	Minor   int
	Major   int
	Patch   int
}

var UpdateStatus = updateStatus{
	Major:   3,
	Minor:   2,
	Patch:   1,
	Latest:  0,
	Unknown: -1,
}

type Script struct {
	Name                 string   `json:"name"`
	PrettyName           string   `json:"prettyName"`
	Version              string   `json:"version"`
	LatestVersion        string   `json:"latestVersion"`
	PackageName          string   `json:"packageName"`
	Maintainer           string   `json:"maintainer"`
	Description          string   `json:"description"`
	URL                  string   `json:"url"`
	RuntimeDependencies  []string `json:"runtimeDependencies"`
	BuildDependencies    []string `json:"buildDependencies"`
	OptionalDependencies []string `json:"optionalDependencies"`
	Breaks               []string `json:"breaks"`
	Gives                string   `json:"gives"`
	Replace              []string `json:"replace"`
	Hash                 string   `json:"hash"`
	PPA                  []string `json:"ppa"`
	PacstallDependencies []string `json:"pacstallDependencies"`
	Patch                []string `json:"patch"`
	Repology             []string `json:"repology"`
	RequiredBy           []string `json:"requiredBy"`
	UpdateStatus         int      `json:"updateStatus"` // enum UpdateStatus
}
