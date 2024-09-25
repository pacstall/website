package timeout

import (
	"time"

	"github.com/joomcode/errorx"
)

type result[T interface{}] struct {
	value T
	err   error
}

func Run[T interface{}](timeoutName string, handle func() (T, error), duration time.Duration) (T, error) {
	var zero T
	resultChan := make(chan result[T], 1)

	go func() {
		value, err := handle()
		resultChan <- result[T]{
			value: value,
			err:   err,
		}
		close(resultChan)
	}()

	select {
	case it := <-resultChan:
		return it.value, it.err
	case <-time.After(duration):
		return zero, errorx.TimeoutElapsed.New("operation %v has timed out", timeoutName)
	}
}
