package array

type Mapper[T any, E any] func(T) E
type Filterer[T any] func(*Iterator[T]) bool
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func Asc[T Ordered]() ComparisonFunc[T] {
	return func(a, b T) bool {
		return a < b
	}
}

func Eq[T Ordered]() ComparisonFunc[T] {
	return func(a, b T) bool {
		return a == b
	}
}

func Desc[T Ordered]() ComparisonFunc[T] {
	return func(a, b T) bool {
		return a > b
	}
}

func Is[T Ordered](value T) Filterer[T] {
	return func(it *Iterator[T]) bool {
		return it.Value == value
	}
}

func Not[T Ordered](value T) Filterer[T] {
	return func(it *Iterator[T]) bool {
		return value != it.Value
	}
}
