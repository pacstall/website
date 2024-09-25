package batch_test

import (
	"testing"
	"time"

	"pacstall.dev/webserver/internal/pacnexus/types/pac/parser/parallelism/batch"
	"pacstall.dev/webserver/internal/pacnexus/types/pac/parser/parallelism/channels"
	"pacstall.dev/webserver/pkg/common/array"
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

	if array.Equals(array.SortBy(channels.ToSlice(out), array.Asc[int]()), items, array.Eq[int]()) == false {
		t.Error("expected results to be sorted")
	}
}
