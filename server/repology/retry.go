package repology

import (
	"fmt"
	"time"

	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/pac"
)

func NewSyncer(maxRetries int) func(*pac.Script) error {
	baseTime := time.Second * 5
	multiplier := 0.2
	retries := maxRetries

	return func(script *pac.Script) error {
		defer func() {
			retries = maxRetries
		}()

		for retries > 0 {

			if multiplier <= 0 {
				multiplier = 1
			}

			computedDelay := baseTime * time.Duration(multiplier)

			retries -= 1
			time.Sleep(computedDelay)

			if retries < maxRetries-1 {
				log.Debug.Println("Trying to sync with repology", computedDelay, multiplier)
			}

			if err := Sync(script); err != nil {
				log.Warn.Println("Failed to fetch repology information", err)
				multiplier *= 1.5
				continue
			}

			multiplier *= 0.9
			return nil
		}

		return fmt.Errorf("Failed to fetch repology information after %v retries", maxRetries)
	}
}
