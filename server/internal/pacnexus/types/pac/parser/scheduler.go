package parser

import (
	"time"

	"pacstall.dev/webserver/pkg/common/log"
	"pacstall.dev/webserver/pkg/common/pacsight"
)

func ScheduleRefresh(every time.Duration, pacsightRpc *pacsight.PacsightRpcService) {
	go refresh(every, pacsightRpc)
}

func refresh(every time.Duration, pacsightRpc *pacsight.PacsightRpcService) {
	for {
		err := ParseAll(pacsightRpc)
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
