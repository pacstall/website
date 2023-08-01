package parser

import (
	"time"

	"pacstall.dev/webserver/log"
)

func ScheduleRefresh(every time.Duration) {
	go func() {
		for {
			err := ParseAll()
			if err != nil {
				log.Error("Failed to parse pacscripts: %v", err)
			}

			time.Sleep(every)
		}
	}()
}
