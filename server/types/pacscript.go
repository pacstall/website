package types

type Pacscript struct {
	Name                 string   `json:"name"`
	Version              string   `json:"version"`
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
	RequiredBy           []string `json:"requiredBy"`
}
