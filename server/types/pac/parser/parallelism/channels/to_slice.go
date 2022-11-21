package channels

import "pacstall.dev/webserver/types/list"

func ToSlice[T any](in <-chan T) []T {
	out := make([]T, 0)
	for item := range in {
		out = append(out, item)
	}

	return out
}

func ToList[T any](in <-chan T) list.List[T] {
	return list.From(ToSlice(in))
}
