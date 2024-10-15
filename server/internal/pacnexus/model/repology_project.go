package model

const RepologyProjectTableName = "repology_projects"

type RepologyProject struct {
	Name string `gorm:"primaryKey"`
}

var RepologyProjectColumns = struct {
	Name string
}{
	Name: "name",
}
