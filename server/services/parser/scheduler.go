package parser

import (
	"time"

	"pacstall.dev/webserver/log"
)

func (s *ParserService) ScheduleRefresh(every time.Duration) {
	go s.refresh(every)
}

func (s *ParserService) refresh(every time.Duration) {
	for {
		err := s.ParseAll()
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
