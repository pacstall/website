package expect

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equals(t *testing.T, subject string, expected interface{}, actual interface{}) {
	t.Helper()

	if !cmp.Equal(expected, actual) {
		diff := cmp.Diff(expected, actual)

		t.Errorf("expected equality for %s. got difference %s", subject, diff)
	}
}
