package internal

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
)

func ExportRepologyDatabase(db *gorm.DB) error {
	err := resetTables(db)
	if err != nil {
		return errors.Join(errors.New("failed to reset repology tables"), err)
	}

	it := 1
	lastProjectName := ""
	for {
		delay := makeSecondDelay()
		defer delay.Wait()

		log.Debug("Page %v | Cursor at: %v", it, lastProjectName)
		projectPage, err := getProjectSearch(lastProjectName)
		if err != nil {
			return errors.Join(errors.New("failed to fetch repology project page"), err)
		}

		var projectProviders []model.RepologyProjectProvider
		var projects []model.RepologyProject
		for projectName, apiProjectProvider := range projectPage {
			lastProjectName = projectName
			for _, apiProjectProvider := range apiProjectProvider {
				projectProvider := mapRepologyApiProjectProviderToModel(projectName, apiProjectProvider)

				projectProviders = append(projectProviders, projectProvider)

				project := model.RepologyProject{
					Name: projectName,
				}

				projects = append(projects, project)
			}
		}

		if len(projects) <= 1 {
			break
		}

		err = db.Save(&projects).Error
		if err != nil {
			return errors.Join(errors.New("failed to create repology projects"), err)
		}

		err = db.CreateInBatches(&projectProviders, 90).Error
		if err != nil {
			return errors.Join(errors.New("failed to create repology project providers"), err)
		}

		it++
	}

	return nil
}

func resetTables(db *gorm.DB) error {
	err := db.AutoMigrate(&model.RepologyProject{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.RepologyProjectProvider{})
	if err != nil {
		return err
	}

	err = db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.RepologyProjectProvider{}).Error
	if err != nil {
		return err
	}

	err = db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.RepologyProject{}).Error
	if err != nil {
		return err
	}

	return nil
}

func makeSecondDelay() *sync.WaitGroup {
	var delay sync.WaitGroup
	delay.Add(1)

	go func() {
		defer delay.Done()
		time.Sleep(1 * time.Second)
	}()

	return &delay
}

func mapRepologyApiProjectProviderToModel(projectName string, apiProjectProvider RepologyApiProject) model.RepologyProjectProvider {
	projectProvider := model.RepologyProjectProvider{
		ProjectName:     projectName,
		Repository:      apiProjectProvider.Repository,
		SubRepository:   apiProjectProvider.SubRepository,
		SourceName:      apiProjectProvider.SourceName,
		VisibleName:     apiProjectProvider.VisibleName,
		BinaryName:      apiProjectProvider.BinaryName,
		Version:         apiProjectProvider.Version,
		OriginalVersion: apiProjectProvider.OriginalVersion,
		Status:          apiProjectProvider.Status,
		Summary:         apiProjectProvider.Summary,
	}
	return projectProvider
}
