package batch

import (
	"sync/atomic"
	"time"
)

func Run[T any, E any](batchSize int, items []T, fn func(T) (E, error)) <-chan E {
	out := make(chan E)
	guard := make(chan interface{}, batchSize)

	left := int32(len(items))
	for _, item := range items {
		go func(item T) {
			// Ensure that parsing in done in queues
			guard <- nil
			result, err := fn(item)

			if err == nil {
				out <- result
			}

			atomic.AddInt32(&left, -1)
			<-guard
		}(item)
	}

	go func() {
		for left != 0 {
			time.Sleep(time.Millisecond)
		}

		close(out)
		close(guard)
	}()

	return out
}
