package batch_test

import (
	"testing"
	"time"

	"pacstall.dev/webserver/parallelism/batch"
	"pacstall.dev/webserver/parallelism/channels"
	"pacstall.dev/webserver/types/list"
)

func Test_Batch_Run(t *testing.T) {
	go func() {
		time.Sleep(time.Millisecond * 70) // Ensure that the batch runner has enough time to run
		panic("timeout")
	}()

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	out := batch.Run(2, items, func(item int) (int, error) {
		time.Sleep(10 * time.Millisecond)
		return item, nil
	})

	if channels.ToList(out).SortBy(list.Asc[int]()).Equals(items, list.Eq[int]()) == false {
		t.Error("Expected results to be sorted")
	}
}
