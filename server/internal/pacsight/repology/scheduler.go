package repology

import (
	"bufio"
	"encoding/gob"
	"io/fs"
	"os"
	"path"
	"time"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/internal/pacsight/config"
	"pacstall.dev/webserver/pkg/common/array"
	"pacstall.dev/webserver/pkg/common/log"
	"pacstall.dev/webserver/pkg/common/types"
)

var Packages types.RepologyApiProjectSearchResponse = make(map[string][]types.RepologyApiProject)

func ScheduleRefresh(every time.Duration) {
	log.Info("attempting to read repology cache...")
	cache, err := readCache()
	if err != nil {
		log.Warn("no cache found. this is normal on clean runs: %+v", err)
	} else {
		log.Info("repology cache read successfully. %d projects found", len(cache))
		Packages = cache
	}

	go func() {
		for {
			log.Info("refreshing Repology database...")
			out := types.RepologyApiProjectSearchResponse{}
			resultsChan, errChan := ExportRepologyDatabase()

		chanLoop:
			for {
				select {
				case pair, ok := <-resultsChan:
					if !ok {
						log.Info("repology database refreshed successfully. %d projects found", len(out))
						break chanLoop
					}

					out[pair.ProjectName] = pair.Projects

					if err := cachePair(pair.ProjectName, pair.Projects); err != nil {
						log.Error("failed to cache Repology projects: %+v", err)
					} else {
						log.Trace("cached repology project '%s'", pair.ProjectName)
					}
				case err := <-errChan:
					log.Error("failed to export Repology projects: %+v", err)
					break chanLoop
				}
			}

			time.Sleep(every)
		}
	}()

	log.Info("scheduled repology refresh every %v", every)
}

func readCache() (types.RepologyApiProjectSearchResponse, error) {
	ensureCacheDirectoryExists()

	var err error
	var dirEntries []fs.DirEntry
	if dirEntries, err = os.ReadDir(config.Repology.CachePath); err != nil {
		return nil, err
	}

	dirEntries = array.Filter(dirEntries, func(entry *array.Iterator[fs.DirEntry]) bool {
		return !entry.Value.IsDir()
	})

	if len(dirEntries) == 0 {
		return nil, errorx.DataUnavailable.New("no cache files found")
	}

	files := array.SwitchMap(dirEntries, func(entry *array.Iterator[fs.DirEntry]) string {
		return entry.Value.Name()
	})

	cache := types.RepologyApiProjectSearchResponse{}

	for _, fileName := range files {
		cacheFilePath := path.Join(config.Repology.CachePath, fileName)
		file, err := os.Open(cacheFilePath)
		if err != nil {
			return nil, err
		}

		defer file.Close()

		reader := bufio.NewReader(file)
		decoder := gob.NewDecoder(reader)

		projects := []types.RepologyApiProject{}
		if err = decoder.Decode(&projects); err != nil {
			return nil, err
		}

		projectName := fileName[:len(fileName)-4] // remove .gob extension

		cache[projectName] = projects
	}

	return cache, nil
}

func cachePair(projectName string, projects []types.RepologyApiProject) error {
	ensureCacheDirectoryExists()
	fileName := projectName + ".gob"
	cacheFilePath := path.Join(config.Repology.CachePath, fileName)

	// delete existing cache file if it exists
	os.RemoveAll(cacheFilePath)

	file, err := os.Create(cacheFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	encoder := gob.NewEncoder(writer)

	if err = encoder.Encode(projects); err != nil {
		return err
	}

	if err = writer.Flush(); err != nil {
		return err
	}

	return nil
}

func ensureCacheDirectoryExists() {
	if err := os.MkdirAll(config.Repology.CachePath, 0755); err != nil {
		log.Fatal("failed to create cache directory: %+v", err)
	}
}
