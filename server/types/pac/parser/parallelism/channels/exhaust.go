package channels

func Exhaust[T any](in <-chan T) {
	for range in {
	}
}
