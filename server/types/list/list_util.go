package list

import (
	"fmt"
	"sync/atomic"
)

type List[T any] []T

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type ComparableList[T Ordered] struct {
	List[T]
}

func New[T any]() List[T] {
	return From(make([]T, 0))
}

func From[T any](items []T) List[T] {
	return items
}

type ComparisonFunc[T any] func(a, b T) bool

func (list List[T]) IndexOf(filter Filterer[T]) (int, error) {
	for idx, it := range list.ToSlice() {
		if filter(it) {
			return idx, nil
		}
	}

	return -1, fmt.Errorf("object does not exist in list")
}

func shallowEqual[T any](a, b T) bool {
	return &a == &b
}

func (list List[T]) Equals(other List[T], compare ComparisonFunc[T]) bool {
	if (list == nil) != (other == nil) {
		return false
	}

	if len(list) != len(other) {
		return false
	}

	for idx, it := range list.ToSlice() {
		if !compare(it, other[idx]) {
			return false
		}
	}

	return true
}

func (list List[T]) Distinct(isEq ComparisonFunc[T]) (out List[T]) {
	out = New[T]()
	for _, item := range list.ToSlice() {
		exists := out.Contains(func(it T) bool { return isEq(it, item) })
		if !exists {
			out = append(out, item)
		}
	}

	return out
}

func (list List[T]) Append(item T) List[T] {
	clone := list.Clone()
	clone = append(clone, item)
	return clone
}

func (list List[T]) Remove(item *T) List[T] {
	clone := list.RemoveFunc(func(a T) bool {
		return shallowEqual(a, *item)
	})

	return clone
}

func (list ComparableList[T]) Remove(item T) List[T] {
	clone := list.RemoveFunc(func(a T) bool {
		return a == item
	})

	return clone
}

func (list List[T]) RemoveFunc(matches func(T) bool) List[T] {
	clone := list.FilterIndex(func(_ int, t T) bool {
		return !matches(t)
	})

	return clone
}

func (list List[T]) IsEmpty() bool {
	return list.Len() == 0
}

func (list List[T]) Find(filter Filterer[T]) (T, error) {
	idx, err := list.IndexOf(filter)
	if err != nil {
		var out T
		return out, err
	}

	return list[idx], nil
}

func (list List[T]) Last() (T, error) {
	if list.IsEmpty() {
		var out T
		return out, fmt.Errorf("list is empty")
	}

	return list[list.Len()-1], nil
}

func (list List[T]) MapIndex(mapper func(int, T) T) List[T] {
	out := make(List[T], list.Len())
	for idx, item := range list {
		value := mapper(idx, item)
		out[idx] = value
	}

	return out
}

func (list List[T]) Map(mapper func(T) T) List[T] {
	return list.MapIndex(func(i int, t T) T {
		return mapper(t)
	})
}

func (list List[T]) MapExt(mapper func(T, List[T]) T) List[T] {
	clone := list.Clone()
	return list.MapIndex(func(i int, t T) T {
		return mapper(t, clone)
	})
}

func (list List[T]) Apply(mapper func([]T) []T) List[T] {
	return mapper(list.Clone())
}

func Apply[T any, E any](list List[T], mapper func([]T) []E) List[E] {
	return mapper(list.Clone())
}

func MapIndex[T any, E any](list List[T], mapper func(int, T) E) List[E] {
	out := make(List[E], list.Len())
	for idx, item := range list {
		value := mapper(idx, item)
		out[idx] = value
	}

	return out
}

func Map[T any, E any](list List[T], mapper func(int, T) E) List[E] {
	return MapIndex(list, func(i int, t T) E {
		return mapper(i, t)
	})
}

func ReduceIndex[T any, E any](list List[T], reducer func(int, T, E) E, accumulator E) E {
	out := accumulator
	for idx, item := range list {
		accumulator = reducer(idx, item, accumulator)
	}

	return out
}

func Reduce[T any, E any](list List[T], reducer func(T, E) E, accumulator E) E {
	return ReduceIndex(list, func(i int, t T, e E) E {
		return reducer(t, e)
	}, accumulator)
}

func (list List[T]) Reverse() List[T] {
	clone := list.Clone()
	for i := clone.Len() - 1; i >= 0; i-- {
		clone = append(clone, clone[i])
	}

	return clone
}

func (list List[T]) FindBy(predicate func(T) bool) (T, error) {
	for _, it := range list.ToSlice() {
		if predicate(it) {
			return it, nil
		}
	}

	var out T
	return out, fmt.Errorf("object does not exist in list")
}

func (list List[T]) Contains(filter Filterer[T]) bool {
	_, err := list.IndexOf(filter)
	return err == nil
}

func (list List[T]) ContainsPtr(item *T) bool {
	_, err := list.IndexOf(func(t T) bool {
		return &t == item
	})
	return err == nil
}

func (list ComparableList[T]) Contains(item T) bool {
	_, err := list.IndexOf(Is(item))
	return err == nil
}

func (list List[T]) Len() int {
	return len(list.ToSlice())
}

func (list List[T]) ToSlice() []T {
	return []T(list)
}

func (list List[T]) SortBy(isLessThan func(T, T) bool) List[T] {
	a := list.Clone()
	return quicksort(a, isLessThan)
}

func quicksort[T any](a []T, isLessThan func(T, T) bool) List[T] {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := len(a) / 2

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if isLessThan(a[i], a[right]) {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	quicksort(a[:left], isLessThan)
	quicksort(a[left+1:], isLessThan)

	return a
}

func (list List[T]) IsSorted(isLessThan func(T, T) bool) bool {
	for i := 1; i < list.Len(); i++ {
		if !isLessThan(list[i-1], list[i]) {
			return false
		}
	}

	return true
}

func (list List[T]) Clone() List[T] {
	return append(List[T](make([]T, 0)), list...)
}

func (list List[T]) FilterIndex(predicate func(int, T) bool) List[T] {
	out := make([]T, 0)
	for idx, it := range list.ToSlice() {
		if predicate(idx, it) {
			out = append(out, it)
		}
	}

	return out
}

func (list List[T]) Filter(predicate func(T) bool) List[T] {
	return list.FilterIndex(func(i int, t T) bool {
		return predicate(t)
	})
}

func (list List[T]) All(predicate func(T) bool) bool {
	for _, it := range list.ToSlice() {
		if !predicate(it) {
			return false
		}
	}

	return true
}

func (list List[T]) Any(predicate func(T) bool) bool {
	passes := false
	for _, it := range list.ToSlice() {
		if predicate(it) {
			passes = true
		}
	}

	return passes
}

func (list List[T]) ToBufChan(ch chan T) chan T {
	left := int32(list.Len())
	for _, item := range list.ToSlice() {
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

func (list List[T]) ToChan() (ch chan T) {
	return list.ToBufChan(ch)
}
