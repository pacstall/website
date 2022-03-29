package channels

import "pacstall.dev/webserver/types/list"

func ToSlice[T any](in <-chan T) []T {
	out := make([]T, 0)
	println("ToSlice: in:", in)
	for item := range in {
		println("ToSlice: item:", item)
		out = append(out, item)
	}

	println("ToSlice: out:", out)
	return out
}

func ToList[T any](in <-chan T) list.List[T] {
	return list.From(ToSlice(in))
}
