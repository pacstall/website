package repology

import (
	"time"

	"pacstall.dev/webserver/log"
)

func (s *RepologyService) ScheduleRefresh(every time.Duration) {
	go func() {

		for {
			log.Info("refreshing Repology database...")
			err := s.ExportRepologyDatabase()
			if err != nil {
				log.Error("failed to export Repology projects: %v", err)
			} else {
				log.Info("repology database refreshed successfully")
			}
			time.Sleep(every)
		}
	}()
}
