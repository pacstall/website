package parser

import (
	"time"

	"pacstall.dev/webserver/log"
)

func ScheduleRefresh(every time.Duration) {
	go refresh(every)
}

func refresh(every time.Duration) {
	for {
		err := ParseAll()
		if err != nil {
			log.Error("parse error: %+v", err)

			retryIn := time.Second * 30
			log.Info("retrying in %v", retryIn)
			time.Sleep(retryIn)
			continue
		}

		time.Sleep(every)
	}
}
