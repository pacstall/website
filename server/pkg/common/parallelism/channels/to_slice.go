package channels

func ToSlice[T any](in <-chan T) []T {
	out := make([]T, 0)
	for item := range in {
		out = append(out, item)
	}

	return out
}
