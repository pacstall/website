package channels_test

import (
	"testing"
	"time"

	"pacstall.dev/webserver/types/pac/parser/parallelism/channels"
)

func registerTimeout() {
	time.Sleep(time.Millisecond * 10)
	panic("timeout")
}

func Test_ToSlice_Empty(t *testing.T) {
	go registerTimeout()

	in := make(chan int)
	close(in)
	out := channels.ToSlice(in)
	if len(out) != 0 {
		t.Errorf("Expected 0, got %d", len(out))
	}
}

func Test_ToSlice(t *testing.T) {
	go registerTimeout()

	in := make(chan int)
	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()
	out := channels.ToSlice(in)

	if len(out) != 3 {
		t.Errorf("Expected 3, got %d", len(out))
	}
}
