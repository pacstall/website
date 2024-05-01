package repository

type RepologyProject interface {
	GetName() string
	SetName(string)
}

type RepologyProjectRepository interface {
	Migrate() error
	Truncate() error
	Save(entities []RepologyProject) error
}
