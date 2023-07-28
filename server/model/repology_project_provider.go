package model

type RepologyProjectProvider struct {
	ID              uint `gorm:"primarykey"`
	ProjectName     string
	Project         RepologyProject `gorm:"foreignKey:Name"`
	Repository      string
	SubRepository   *string `gorm:"default:null"`
	SourceName      *string `gorm:"default:null"`
	VisibleName     *string `gorm:"default:null"`
	BinaryName      *string `gorm:"default:null"`
	Version         string
	OriginalVersion string
	Status          string
	Summary         string
}

var RepologyProjectProviderColumns = struct {
	ID              string
	ProjectName     string
	Repository      string
	SubRepository   string
	SourceName      string
	VisibleName     string
	BinaryName      string
	Version         string
	OriginalVersion string
	Status          string
	Summary         string
}{
	ID:              "id",
	ProjectName:     "project_name",
	Repository:      "repository",
	SubRepository:   "sub_repository",
	SourceName:      "source_name",
	VisibleName:     "visible_name",
	BinaryName:      "binary_name",
	Version:         "version",
	OriginalVersion: "original_version",
	Status:          "status",
	Summary:         "summary",
}
