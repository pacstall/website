package repology

const (
	repologProjectUrl = "https://repology.org/api/v1/project/%s"
)

type repologyProject struct {
	PrettyName string
	Version    string
}

type repologyRawProject = map[string]interface{}

type repologySemiRawProject struct {
	Version     string
	VisibleName string
	Name        string
	Status      string
	SrcName     string
	BinName     string
	Repo        string
	SubRepo     string
	Licenses    []string
	OrigVersion string
	Summary     string
	Maintainers []string
	Categories  []string
}
