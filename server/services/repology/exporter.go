package repology

import (
	"strings"
	"time"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
	"pacstall.dev/webserver/services/repology/internal"
	"pacstall.dev/webserver/types/repository"
)

func (s *RepologyService) ExportRepologyDatabase() error {
	if err := s.migrateTables(); err != nil {
		return errorx.Decorate(err, "failed to reset repology tables")
	}

	it := 1
	lastProjectName := ""

	const REPOLOGY_PROJECT_FETCH_THROTTLE = 400 * time.Millisecond

	lastRepoFetch := time.Now()

	for {
		if time.Since(lastRepoFetch) < REPOLOGY_PROJECT_FETCH_THROTTLE {
			time.Sleep(REPOLOGY_PROJECT_FETCH_THROTTLE - time.Since(lastRepoFetch))
		}

		log.Debug("page %v | cursor at: %v", it, lastProjectName)
		projectPage, err := internal.GetProjectSearch(lastProjectName)
		if err != nil {
			return errorx.Decorate(err, "failed to fetch repology project page")
		}

		lastRepoFetch = time.Now()

		var projects []repository.RepologyProject
		var projectProviders []repository.RepologyProjectProvider
		for projectName, apiProjectProvider := range projectPage {

			lastProjectName = identityOrSkipProject(projectName)
			for _, apiProjectProvider := range apiProjectProvider {
				// Save project provider as inactive
				projectProvider := mapRepologyApiProjectProviderToModel(projectName, apiProjectProvider)
				projectProviders = append(projectProviders, &projectProvider)

				project := model.RepologyProject{
					Name: projectName,
				}

				projects = append(projects, &project)
			}
		}

		if len(projects) <= 1 {
			break
		}

		if err = s.repologyProjectRepository.Save(projects); err != nil {
			return err
		}

		if err = s.repologyProjectProviderRepository.CreateInBatches(projectProviders, 90); err != nil {
			return err
		}

		it++
	}

	// Delete active (old) repology project providers
	if err := s.repologyProjectProviderRepository.DeleteWhereActive(); err != nil {
		return err
	}

	// Mark new repology project providers as active
	if err := s.repologyProjectProviderRepository.UpdateAllAsActive(); err != nil {
		return err
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

func (s *RepologyService) migrateTables() error {
	if err := s.repologyProjectRepository.Migrate(); err != nil {
		return err
	}

	if err := s.repologyProjectRepository.Truncate(); err != nil {
		return err
	}

	if err := s.repologyProjectProviderRepository.Migrate(); err != nil {
		return err
	}

	if err := s.repologyProjectProviderRepository.Truncate(); err != nil {
		return err
	}

	return nil
}

func mapRepologyApiProjectProviderToModel(projectName string, apiProjectProvider internal.RepologyApiProject) model.RepologyProjectProvider {
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
