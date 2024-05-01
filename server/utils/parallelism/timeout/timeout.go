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
	errChan := make(chan error)
	resultChan := make(chan result[T])

	go func() {
		value, err := handle()
		resultChan <- result[T]{
			value: value,
			err:   err,
		}
	}()

	go func() {
		time.Sleep(duration)
		errChan <- errorx.TimeoutElapsed.New("operation %v has timed out", timeoutName)
	}()

	select {
	case it := <-resultChan:
		return it.value, it.err
	case err := <-errChan:
		return zero, err
	}
}
