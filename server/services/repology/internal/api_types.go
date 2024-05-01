package internal

type RepologyApiProject struct {
	Repository      string   `json:"repo"`
	SubRepository   *string  `json:"subrepo"`
	SourceName      *string  `json:"srcname"`
	VisibleName     *string  `json:"visiblename"`
	BinaryName      *string  `json:"binname"`
	Version         string   `json:"version"`
	OriginalVersion string   `json:"origversion"`
	Status          string   `json:"status"`
	Summary         string   `json:"summary"`
	Licenses        []string `json:"licenses"`
	Maintainers     []string `json:"maintainers"`
	Categories      []string `json:"categories"`
}

type RepologyApiProjectSearchResponse = map[string][]RepologyApiProject
