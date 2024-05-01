package service

import (
	"time"

	"pacstall.dev/webserver/types/pac"
)

type RepologyService interface {
	ScheduleRefresh(time time.Duration)
	Sync(script *pac.Script) error
}
