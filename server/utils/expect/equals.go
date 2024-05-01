package expect

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equals[T interface{}](t *testing.T, subject string, expected T, actual T) {
	t.Helper()

	if !cmp.Equal(expected, actual) {
		diff := cmp.Diff(expected, actual)

		t.Errorf("expected equality for %s. got difference %s", subject, diff)
	}
}

func False(t *testing.T, message string, condition bool) {
	t.Helper()

	if condition {
		t.Errorf(message)
	}
}

func True(t *testing.T, message string, condition bool) {
	t.Helper()

	if !condition {
		t.Errorf(message)
	}
}
