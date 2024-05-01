package model

import (
	"fmt"

	"github.com/joomcode/errorx"
	"gorm.io/gorm"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/repository"
)

const RepologyProjectProviderTableName = "repology_project_providers"

type RepologyProjectProvider struct {
	ID              uint                       `gorm:"primarykey"`
	ProjectName     string                     `gorm:"index:"`
	Project         repository.RepologyProject `gorm:"foreignKey:Name"`
	Repository      string
	SubRepository   *string `gorm:"default:null"`
	SourceName      *string `gorm:"index:,default:null"`
	VisibleName     *string `gorm:"index:,default:null"`
	BinaryName      *string `gorm:"index:,default:null"`
	Version         string
	OriginalVersion string
	Status          string
	Summary         string
	Active          bool `gorm:"index:,default:false"`
}

func (r *RepologyProjectProvider) GetID() uint {
	return r.ID
}
func (r *RepologyProjectProvider) SetID(ID uint) {
	r.ID = ID
}

func (r *RepologyProjectProvider) GetProjectName() string {
	return r.ProjectName

}
func (r *RepologyProjectProvider) SetProjectName(ProjectName string) {
	r.ProjectName = ProjectName
}

func (r *RepologyProjectProvider) GetProject() repository.RepologyProject {
	return r.Project

}
func (r *RepologyProjectProvider) SetProject(Project repository.RepologyProject) {
	r.Project = Project
}
func (r *RepologyProjectProvider) GetRepository() string {
	return r.Repository

}
func (r *RepologyProjectProvider) SetRepository(Repository string) {
	r.Repository = Repository
}
func (r *RepologyProjectProvider) GetSubRepository() *string {
	return r.SubRepository

}
func (r *RepologyProjectProvider) SetSubRepository(SubRepository *string) {
	r.SubRepository = SubRepository
}
func (r *RepologyProjectProvider) GetSourceName() *string {
	return r.SourceName

}
func (r *RepologyProjectProvider) SetSourceName(SourceName *string) {
	r.SourceName = SourceName
}
func (r *RepologyProjectProvider) GetVisibleName() *string {
	return r.VisibleName

}
func (r *RepologyProjectProvider) SetVisibleName(VisibleName *string) {
	r.VisibleName = VisibleName
}
func (r *RepologyProjectProvider) GetBinaryName() *string {
	return r.BinaryName

}
func (r *RepologyProjectProvider) SetBinaryName(BinaryName *string) {
	r.BinaryName = BinaryName
}
func (r *RepologyProjectProvider) GetVersion() string {
	return r.Version

}
func (r *RepologyProjectProvider) SetVersion(Version string) {
	r.Version = Version
}
func (r *RepologyProjectProvider) GetOriginalVersion() string {
	return r.OriginalVersion

}
func (r *RepologyProjectProvider) SetOriginalVersion(OriginalVersion string) {
	r.OriginalVersion = OriginalVersion
}
func (r *RepologyProjectProvider) GetStatus() string {
	return r.Status

}
func (r *RepologyProjectProvider) SetStatus(Status string) {
	r.Status = Status
}
func (r *RepologyProjectProvider) GetSummary() string {
	return r.Summary

}
func (r *RepologyProjectProvider) SetSummary(Summary string) {
	r.Summary = Summary
}
func (r *RepologyProjectProvider) GetActive() bool {
	return r.Active

}
func (r *RepologyProjectProvider) SetActive(Active bool) {
	r.Active = Active
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
	Active          string
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
	Active:          "active",
}

type RepologyProjectProviderRepository struct {
	database *gorm.DB
}

func InitRepologyProjectProviderRepository(db *gorm.DB) *RepologyProjectProviderRepository {
	r := &RepologyProjectProviderRepository{}

	r.database = db

	return r
}

func (r *RepologyProjectProviderRepository) CreateInBatches(entities []repository.RepologyProjectProvider, batchSize int) error {
	err := r.database.CreateInBatches(entities, batchSize).Error
	if err != nil {
		return errorx.Decorate(err, "failed to create repology project providers")
	}

	return nil
}

func (r *RepologyProjectProviderRepository) DeleteWhereActive() error {
	if err := r.database.Debug().Where(fmt.Sprintf("%v = ?", RepologyProjectProviderColumns.Active), true).Delete(&RepologyProjectProvider{}).Error; err != nil {
		return errorx.Decorate(err, "failed to delete old repology project providers")
	}

	return nil
}

func (r *RepologyProjectProviderRepository) UpdateAllAsActive() error {
	if err := r.database.Debug().Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s = 1",
			RepologyProjectProviderTableName,
			RepologyProjectProviderColumns.Active,
		),
	).Error; err != nil {
		return errorx.Decorate(err, "failed to update new repology project providers")
	}

	return nil
}

func (r *RepologyProjectProviderRepository) Migrate() error {
	return r.database.AutoMigrate(&RepologyProjectProvider{})
}
func (r *RepologyProjectProviderRepository) Truncate() error {
	log.Debug("attempting to truncate table %v", RepologyProjectProviderTableName)
	err := r.database.Exec("TRUNCATE TABLE " + RepologyProjectProviderTableName).Error
	if err != nil {
		return errorx.Decorate(err, "failed to truncate table %v", RepologyProjectProviderTableName)
	}

	log.Info("successfully truncated table %v", RepologyProjectProviderTableName)
	return nil
}

func (r *RepologyProjectProviderRepository) FindAllWhereProjectNameAndFiltersSortedByVersionDesc(
	projectName string,
	filters map[string]interface{},
) ([]repository.RepologyProjectProvider, error) {
	query := r.database.Where(
		fmt.Sprintf("%v = ?", RepologyProjectProviderColumns.ProjectName),
		projectName,
	).Where(
		fmt.Sprintf("%v = ?", RepologyProjectProviderColumns.Active),
		true,
	)

	for column, value := range filters {
		query = query.Where(fmt.Sprintf("%v = ?", column), value)
	}

	results := []*RepologyProjectProvider{}
	if err := query.Order(fmt.Sprintf("%v desc", RepologyProjectProviderColumns.Version)).Find(&results).Error; err != nil || len(results) == 0 {
		return nil, errorx.Decorate(err, "failed to fetch repology project")
	}

	// Data is not copied as they're reference types
	mappedResults := make([]repository.RepologyProjectProvider, len(results))
	for idx, value := range results {
		mappedResults[idx] = value
	}

	return mappedResults, nil
}
