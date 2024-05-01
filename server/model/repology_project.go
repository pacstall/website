package model

import (
	"github.com/joomcode/errorx"
	"gorm.io/gorm"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/repository"
)

const RepologyProjectTableName = "repology_projects"

type RepologyProject struct {
	Name string `gorm:"primaryKey"`
}

func (r *RepologyProject) GetName() string {
	return r.Name
}

func (r *RepologyProject) SetName(name string) {
	r.Name = name
}

var RepologyProjectColumns = struct {
	Name string
}{
	Name: "name",
}

type RepologyProjectRepository struct {
	database *gorm.DB
}

func InitRepologyProjectRepository(r *RepologyProjectRepository, db *gorm.DB) *RepologyProjectRepository {
	if r == nil {
		r = &RepologyProjectRepository{}
	}

	r.database = db

	return r
}

func (r *RepologyProjectRepository) Save(projects []repository.RepologyProject) error {
	err := r.database.Save(&projects).Error
	if err != nil {
		return errorx.Decorate(err, "failed to save repology projects")
	}

	return nil
}

func (r *RepologyProjectRepository) Migrate() error {
	return r.database.AutoMigrate(&RepologyProjectProvider{})
}

func (r *RepologyProjectRepository) Truncate() error {
	log.Debug("attempting to truncate table %v", RepologyProjectTableName)
	err := r.database.Exec("TRUNCATE TABLE " + RepologyProjectTableName).Error
	if err != nil {
		return errorx.Decorate(err, "failed to truncate table %v", RepologyProjectTableName)
	}

	log.Info("successfully truncated table %v", RepologyProjectTableName)
	return nil
}
