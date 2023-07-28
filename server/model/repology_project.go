package model

type RepologyProject struct {
	Name string `gorm:"primaryKey"`
}

var RepologyProjectColumns = struct {
	Name string
}{
	Name: "name",
}
