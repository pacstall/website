package list

// ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string

type Mapper[T any, E any] func(T) E
type Filterer[T any] func(T) bool

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

func AscBy[T any, E Ordered](mapper Mapper[T, E]) ComparisonFunc[T] {
	return func(a, b T) bool {
		return mapper(a) < mapper(b)
	}
}

func DescBy[T any, E Ordered](mapper Mapper[T, E]) ComparisonFunc[T] {
	return func(a, b T) bool {
		return mapper(a) > mapper(b)
	}
}

func Is[T Ordered](value T) Filterer[T] {
	return func(a T) bool {
		return a == value
	}
}

func Not[T Ordered](value T) Filterer[T] {
	return func(a T) bool {
		return a != value
	}
}

func IsBy[T any, E Ordered](value E, mapper Mapper[T, E]) Filterer[T] {
	return func(a T) bool {
		return mapper(a) == value
	}
}

func NotBy[T any, E Ordered](value E, mapper Mapper[T, E]) Filterer[T] {
	return func(a T) bool {
		return mapper(a) != value
	}
}
