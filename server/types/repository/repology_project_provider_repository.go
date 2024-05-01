package repository

type RepologyProjectProvider interface {
	GetID() uint
	SetID(uint)

	GetProjectName() string
	SetProjectName(string)

	GetProject() RepologyProject
	SetProject(RepologyProject)

	GetRepository() string
	SetRepository(string)

	GetSubRepository() *string
	SetSubRepository(*string)

	GetSourceName() *string
	SetSourceName(*string)

	GetVisibleName() *string
	SetVisibleName(*string)

	GetBinaryName() *string
	SetBinaryName(*string)

	GetVersion() string
	SetVersion(string)

	GetOriginalVersion() string
	SetOriginalVersion(string)

	GetStatus() string
	SetStatus(string)

	GetSummary() string
	SetSummary(string)

	GetActive() bool
	SetActive(bool)
}

type RepologyProjectProviderRepository interface {
	Migrate() error
	Truncate() error
	CreateInBatches(entities []RepologyProjectProvider, batchSize int) error
	DeleteWhereActive() error
	UpdateAllAsActive() error
	FindAllWhereProjectNameAndFiltersSortedByVersionDesc(projectName string, filters map[string]interface{}) ([]RepologyProjectProvider, error)
}
