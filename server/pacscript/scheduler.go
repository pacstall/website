package pacscript

import "time"

func ScheduleRefresh(every time.Duration) {
	go func() {
		for {
			time.Sleep(every)
			Load()
		}
	}()
}
