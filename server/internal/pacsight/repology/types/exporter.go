package types

import (
	"math"
	"strings"
	"time"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/pkg/common/log"
	"pacstall.dev/webserver/pkg/common/types"
)

// Ends up waiting `sum(1^2, 2^2, 3^2, ..., RETRY_COUNT^2)` = 385 seconds (6.4 minutes) at most
const RETRY_COUNT = 10
const REPOLOGY_PROJECT_FETCH_THROTTLE = time.Second

func ExportRepologyDatabase() (types.RepologyApiProjectSearchResponse, error) {
	page := 1
	lastProjectName := ""

	lastRepoFetch := time.Now()
	var repologyProjectsMap = make(types.RepologyApiProjectSearchResponse)

	for {
		if time.Since(lastRepoFetch) < REPOLOGY_PROJECT_FETCH_THROTTLE {
			time.Sleep(REPOLOGY_PROJECT_FETCH_THROTTLE - time.Since(lastRepoFetch))
		}

		log.Debug("page %v | cursor at: %v", page, lastProjectName)

		var projectPage map[string][]types.RepologyApiProject
		var err error

	retry:
		for i := 1; i <= RETRY_COUNT; i += 1 {
			projectPage, err = getProjectSearch(lastProjectName)
			if err == nil {
				break retry
			}

			retryDelay := time.Duration(math.Pow(float64(i), 2)) * REPOLOGY_PROJECT_FETCH_THROTTLE
			log.Debug("failed to fetch repology project page '%s'. retrying in %v", lastProjectName, retryDelay)
			time.Sleep(retryDelay)
		}

		if err != nil {
			return nil, errorx.Decorate(err, "failed to fetch repology project page '%s'", lastProjectName)
		}

		lastRepoFetch = time.Now()

		shouldStop := false
		for projectName, apiProjectProvider := range projectPage {
			if _, ok := repologyProjectsMap[projectName]; !ok {
				repologyProjectsMap[projectName] = []types.RepologyApiProject{}
			}

			repologyProjectsMap[projectName] = append(repologyProjectsMap[projectName], apiProjectProvider...)

			shouldStop = projectName == lastProjectName
			lastProjectName = identityOrSkipProject(projectName)
		}

		if shouldStop {
			break
		}

		page += 1
	}

	return repologyProjectsMap, nil
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
