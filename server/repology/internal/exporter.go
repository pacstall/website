package internal

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
)

func ExportRepologyDatabase(db *gorm.DB) error {
	err := migrateTables(db)
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
				// Save project provider as inactive
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

	// Delete active (old) repology project providers
	if err := db.Where(fmt.Sprintf("%v = ?", model.RepologyProjectProviderColumns.Active), true).Delete(&model.RepologyProjectProvider{}).Error; err != nil {
		return errors.Join(errors.New("failed to delete old repology project providers"), err)
	}

	// Mark new repology project providers as active
	if err := db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s = 1",
			model.RepologyProjectProviderTableName,
			model.RepologyProjectProviderColumns.Active,
		),
	).Error; err != nil {
		return errors.Join(errors.New("failed to update new repology project providers"), err)
	}

	return nil
}

func migrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(&model.RepologyProject{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.RepologyProjectProvider{})
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
		// Wait 750ms before making another request
		// Repology API has a rate limit of 1 request per second but some requests take longer than 1 second so it averages out
		time.Sleep(750 * time.Millisecond)
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
		Active:          false,
	}
	return projectProvider
}
