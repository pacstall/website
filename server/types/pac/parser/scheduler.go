package parser

import (
	"time"

	"pacstall.dev/webserver/log"
)

func ScheduleRefresh(every time.Duration) {
	go func() {
		for {
			time.Sleep(every)
			err := ParseAll()
			if err != nil {
				log.Error("Failed to parse pacscripts: %v", err)
			}

		}
	}()
}
