package service

import "time"

type ParserService interface {
	ScheduleRefresh(interval time.Duration)
}
