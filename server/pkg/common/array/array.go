package array

import (
	"sort"
	"sync/atomic"

	"github.com/joomcode/errorx"
)

type ComparisonFunc[T any] func(a, b T) bool

func SortBy[T any](arr []T, isLessThan func(T, T) bool) []T {
	sort.SliceStable(arr, func(i, j int) bool {
		return isLessThan(arr[i], arr[j])
	})

	return arr
}

func IndexOf[T any](arr []T, filter Filterer[T]) (int, error) {
	for idx, it := range arr {
		if filter(&Iterator[T]{idx, it, arr}) {
			return idx, nil
		}
	}

	return -1, errorx.IllegalState.New("object does not exist in list")
}

func Find[T any](arr []T, filter Filterer[T]) (T, error) {
	idx, err := IndexOf(arr, filter)
	if err != nil {
		var out T
		return out, err
	}

	return arr[idx], nil
}

func Equals[T any](arr []T, other []T, compare ComparisonFunc[T]) bool {
	if (arr == nil) != (other == nil) {
		return false
	}

	if len(arr) != len(other) {
		return false
	}

	for idx, it := range arr {
		if !compare(it, other[idx]) {
			return false
		}
	}

	return true
}

func Distinct[T any](arr []T, isEq ComparisonFunc[T]) []T {
	out := make([]T, 0)
	for _, item := range arr {
		exists := Contains(out, func(it *Iterator[T]) bool { return isEq(it.Value, item) })
		if !exists {
			out = append(out, item)
		}
	}

	return out
}

func IsEmpty[T any](arr []T) bool {
	return len(arr) == 0
}

func FindBy[T any](arr []T, predicate func(T) bool) (T, error) {
	for _, it := range arr {
		if predicate(it) {
			return it, nil
		}
	}

	var out T
	return out, errorx.IllegalState.New("object does not exist in list")
}

func Contains[T any](arr []T, filter Filterer[T]) bool {
	_, err := IndexOf(arr, filter)
	return err == nil
}

func ContainsPtr[T any](arr []T, item *T) bool {
	_, err := IndexOf(arr, func(it *Iterator[T]) bool {
		return &it.Value == item
	})
	return err == nil
}

func IsSorted[T any](arr []T, isLessThan func(T, T) bool) bool {
	for i := 1; i < len(arr); i++ {
		if !isLessThan(arr[i-1], arr[i]) {
			return false
		}
	}

	return true
}

func Clone[T any](arr []T) []T {
	return append(make([]T, 0), arr...)
}

func Filter[T any](arr []T, predicate func(*Iterator[T]) bool) []T {
	out := make([]T, 0)
	for idx, it := range arr {
		if predicate(&Iterator[T]{idx, it, arr}) {
			out = append(out, it)
		}
	}

	return out
}

func FilterPtr[T any](arr []T, predicate func(*PtrIterator[T]) bool) []T {
	out := make([]T, 0)
	for idx, it := range arr {
		if predicate(&PtrIterator[T]{idx, &it, arr}) {
			out = append(out, it)
		}
	}

	return out
}

func All[T any](arr []T, predicate func(T) bool) bool {
	for _, it := range arr {
		if !predicate(it) {
			return false
		}
	}

	return true
}

func Any[T any](arr []T, predicate func(T) bool) bool {
	passes := false
	for _, it := range arr {
		if predicate(it) {
			passes = true
		}
	}

	return passes
}

func ToBufChan[T any](arr []T, ch chan T) chan T {
	left := int32(len(arr))
	for _, item := range arr {
		go func(item T) {
			ch <- item
			atomic.AddInt32(&left, -1)
			if left == 0 {
				close(ch)
			}
		}(item)
	}
	return ch
}

func ToChan[T any](arr []T) (ch chan T) {
	return ToBufChan(arr, ch)
}

func Last[T any](arr []T) (T, error) {
	if IsEmpty(arr) {
		var out T
		return out, errorx.IllegalState.New("array is empty")
	}

	return arr[len(arr)-1], nil
}

type Iterator[T any] struct {
	Index int
	Value T
	Array []T
}

type PtrIterator[T any] struct {
	Index int
	Value *T
	Array []T
}

/*  Maps the given array in place. */
func Map[T any](arr []T, mapper func(it *Iterator[T]) T) []T {
	for idx, item := range arr {
		value := mapper(&Iterator[T]{idx, item, arr})
		arr[idx] = value
	}

	return arr
}

func SwitchMap[T any, E any](arr []T, mapper func(it *Iterator[T]) E) []E {
	out := make([]E, len(arr))
	for idx, item := range arr {
		value := mapper(&Iterator[T]{idx, item, arr})
		out[idx] = value
	}

	return out
}

/*  Maps the given array in place. */
func MapPtr[T any](arr []T, mapper func(it *PtrIterator[T]) T) []T {
	for idx, item := range arr {
		value := mapper(&PtrIterator[T]{idx, &item, arr})
		arr[idx] = value
	}

	return arr
}

func SwitchMapPtr[T any, E any](arr []T, mapper func(it *PtrIterator[T]) E) []E {
	out := make([]E, len(arr))
	for idx, item := range arr {
		value := mapper(&PtrIterator[T]{idx, &item, arr})
		out[idx] = value
	}

	return out
}

func ReduceIndex[T any, E any](arr []T, reducer func(int, T, E) E, accumulator E) E {
	out := accumulator
	for idx, item := range arr {
		out = reducer(idx, item, out)
	}

	return out
}

func Reduce[T any, E any](arr []T, reducer func(T, E) E, accumulator E) E {
	return ReduceIndex(arr, func(i int, t T, e E) E {
		return reducer(t, e)
	}, accumulator)
}

/* Reverse the given array in place. */
func Reverse[T any](arr []T) []T {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
