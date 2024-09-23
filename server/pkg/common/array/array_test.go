package array_test

import (
	"testing"

	"pacstall.dev/webserver/types/array"
)

func Test_Array_SortBy(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	array.
		SortBy(data, array.Asc[int]())

	if !array.IsSorted(data, array.Asc[int]()) {
		t.Errorf("Expected array to be sorted")
	}
}

func Test_Array_IsSorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	if !array.IsSorted(data, array.Asc[int]()) {
		t.Errorf("Expected array to be sorted")
	}
}

func Test_Array_IsSorted_Fail(t *testing.T) {
	data := []int{1, 2, 6, 4, 5}

	if array.IsSorted(data, array.Asc[int]()) {
		t.Errorf("Expected array to not be sorted")
	}
}

func Test_Array_IsSorted_Desc(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}

	if !array.IsSorted(data, array.Desc[int]()) {
		t.Errorf("Expected array to be sorted")
	}
}

func Test_Array_IsSorted_Desc_Fail(t *testing.T) {
	data := []int{5, 4, 5, 2, 1}

	if array.IsSorted(data, array.Desc[int]()) {
		t.Errorf("Expected list to not be sorted")
	}
}

func Test_Array_Filter(t *testing.T) {
	data := []int{5, 4, 7, 7, 2, 1}
	data = array.Filter(data, array.Not[int](7))

	if len(data) != 4 {
		t.Errorf("Expected 4, got %d", len(data))
	}
}

func Test_Array_Filter_Inclussive(t *testing.T) {
	data := []int{5, 7, 7, 2, 3, 1}

	data = array.Filter(data, array.Is(7))

	if len(data) != 2 {
		t.Errorf("Expected 2, got %d", len(data))
	}
}

func Test_Array_Contains(t *testing.T) {
	found := array.Contains([]int{5, 4, 7, 7, 2, 1}, array.Is(2))
	if !found {
		t.Errorf("Expected to find 2")
	}
}

func Test_Array_Contains_Fail(t *testing.T) {
	found := array.Contains([]int{5, 4, 7, 7, 2, 1}, array.Is(10))

	if found {
		t.Errorf("Expected to not find 10")
	}
}

func Test_Find(t *testing.T) {
	ten, err := array.Find([]int{5, 4, 7, 7, 2, 10}, array.Is(10))

	if err != nil {
		t.Errorf("Expected to find 10")
	}

	if ten != 10 {
		t.Errorf("Expected 10, got %d", ten)
	}
}

func Test_Find_Fail(t *testing.T) {
	_, err := array.Find([]int{5, 4, 7, 7, 2, 10}, array.Is(11))

	if err == nil {
		t.Errorf("Expected to not find 11")
	}
}

func Test_Distinct(t *testing.T) {
	items := []int{4, 7, 10, 4, 7, 10}

	expected := []int{4, 7, 10}

	actualDistinct := array.Distinct(items, array.Eq[int]())
	if !array.Equals(actualDistinct, expected, array.Eq[int]()) {
		t.Errorf("Expected %v, got %v", expected, items)
	}
}
