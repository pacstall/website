package repology

import (
	"time"

	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
)

func ScheduleRefresh(every time.Duration) {
	go func() {
		for {
			db := model.Instance()
			log.Info("Refreshing Repology database...")
			err := ExportRepologyDatabase(db)
			if err != nil {
				log.Error("Failed to export Repology projects: %v", err)
			} else {
				log.Info("Repology database refreshed successfully")
			}

			time.Sleep(every)
		}
	}()
}
