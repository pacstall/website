package repology

import (
	"encoding/json"
	"os"
	"time"

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
		log.Info("repology cache read successfully")
		Packages = cache
	}

	go func() {
		for {
			log.Info("refreshing Repology database...")
			results, err := ExportRepologyDatabase()
			if err != nil {
				log.Error("failed to export Repology projects: %+v", err)
			} else {
				log.Info("repology database refreshed successfully")
				Packages = results
				if err = cacheValues(Packages); err != nil {
					log.Error("failed to cache Repology projects: %+v", err)
				}
			}
			time.Sleep(every)
		}
	}()

	log.Info("scheduled repology refresh every %v", every)
}

const CACHE_FILE_NAME = "repology_cache.json"

func readCache() (types.RepologyApiProjectSearchResponse, error) {
	bytes, err := os.ReadFile(CACHE_FILE_NAME)
	if err != nil {
		return nil, err
	}

	var cache types.RepologyApiProjectSearchResponse = make(map[string][]types.RepologyApiProject)
	if err = json.Unmarshal(bytes, &cache); err != nil {
		return nil, err
	}

	return cache, nil
}

func cacheValues(value types.RepologyApiProjectSearchResponse) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return os.WriteFile(CACHE_FILE_NAME, bytes, 0644)
}
