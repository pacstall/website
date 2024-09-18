package internal

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/joomcode/errorx"
	"gorm.io/gorm"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
)

const RETRY_COUNT = 5

func ExportRepologyDatabase(db *gorm.DB) error {
	err := migrateTables(db)
	if err != nil {
		return errors.Join(errors.New("failed to reset repology tables"), err)
	}

	it := 1
	lastProjectName := ""

	const REPOLOGY_PROJECT_FETCH_THROTTLE = time.Second

	lastRepoFetch := time.Now()

	for {
		if time.Since(lastRepoFetch) < REPOLOGY_PROJECT_FETCH_THROTTLE {
			time.Sleep(REPOLOGY_PROJECT_FETCH_THROTTLE - time.Since(lastRepoFetch))
		}

		log.Debug("page %v | cursor at: %v", it, lastProjectName)

		var projectPage map[string][]RepologyApiProject
		var err error

	retry:
		for i := 0; i < RETRY_COUNT; i += 1 {
			projectPage, err = getProjectSearch(lastProjectName)
			if err == nil {
				break retry
			}

			retryDelay := time.Duration(i+1) * REPOLOGY_PROJECT_FETCH_THROTTLE
			log.Debug("failed to fetch repology project page '%s'. retrying in %v", lastProjectName, retryDelay)
			time.Sleep(retryDelay)
		}

		if err != nil {
			return errorx.Decorate(err, "failed to fetch repology project page '%s'", lastProjectName)
		}

		lastRepoFetch = time.Now()

		var projects []model.RepologyProject
		var projectProviders []model.RepologyProjectProvider
		for projectName, apiProjectProvider := range projectPage {

			lastProjectName = identityOrSkipProject(projectName)
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
	if err := db.Debug().Where(fmt.Sprintf("%v = ?", model.RepologyProjectProviderColumns.Active), true).Delete(&model.RepologyProjectProvider{}).Error; err != nil {
		return errors.Join(errors.New("failed to delete old repology project providers"), err)
	}

	// Mark new repology project providers as active
	if err := db.Debug().Exec(
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

var projectNamesToSkipToNextCussor = map[string]string{
	"emacs:":   "emacsa",
	"go:":      "goa",
	"haskell:": "haskella",
	"lisp:":    "lispa",
	"node:":    "nodea",
	"ocaml:":   "ocamla",
	"perl:":    "perla",
	"php:":     "phpa",
	"python:":  "pythona",
	"r:":       "ra",
	"ruby:":    "rubya",
	"rust:":    "rusta",
	"texlive:": "texlivea",
}

func identityOrSkipProject(name string) string {
	for prefix, skipTo := range projectNamesToSkipToNextCussor {
		if strings.HasPrefix(name, prefix) {
			return skipTo
		}
	}

	return name
}

func migrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(&model.RepologyProject{})
	if err != nil {
		return err
	}

	if err = truncateTable(db, model.RepologyProjectTableName); err != nil {
		return err
	}

	err = db.AutoMigrate(&model.RepologyProjectProvider{})
	if err != nil {
		return err
	}

	if err = truncateTable(db, model.RepologyProjectProviderTableName); err != nil {
		return err
	}

	return nil
}

func truncateTable(db *gorm.DB, tableName string) error {
	log.Debug("attempting to truncate table %v", tableName)
	err := db.Exec("TRUNCATE TABLE " + tableName).Error
	if err != nil {
		return errors.Join(fmt.Errorf("failed to truncate table %v", tableName), err)
	}

	log.Info("successfully truncated table %v", tableName)
	return nil
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
