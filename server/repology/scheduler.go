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
			log.Info("refreshing Repology database...")
			err := ExportRepologyDatabase(db)
			if err != nil {
				log.Error("failed to export Repology projects: %v", err)
			} else {
				log.Info("repology database refreshed successfully")
			}

			time.Sleep(every)
		}
	}()
}
