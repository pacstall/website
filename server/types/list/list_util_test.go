package list_test

import (
	"testing"

	"pacstall.dev/webserver/types/list"
)

func Test_List_Append_Remove(t *testing.T) {
	data := list.New[int]().
		Append(1).
		Append(2).
		Append(3).
		Append(4).
		Append(5).
		Append(6).
		RemoveFunc(list.Is(6))

	if data.Len() != 5 {
		t.Errorf("Expected 5, got %d", data.Len())
	}

	if data[0] != 1 {
		t.Errorf("Expected 1, got %d", data[0])
	}

	if data[1] != 2 {
		t.Errorf("Expected 2, got %d", data[1])
	}

	if data[2] != 3 {
		t.Errorf("Expected 3, got %d", data[2])
	}

	if data[3] != 4 {
		t.Errorf("Expected 4, got %d", data[3])
	}

	if data[4] != 5 {
		t.Errorf("Expected 5, got %d", data[4])
	}
}

func Test_List_SortBy(t *testing.T) {
	data := list.New[int]().
		Append(5).
		Append(3).
		Append(6).
		Append(2).
		Append(1).
		SortBy(list.Asc[int]())

	if !data.IsSorted(list.Asc[int]()) {
		t.Errorf("Expected list to be sorted")
	}
}

func Test_List_IsSorted(t *testing.T) {
	data := list.New[int]().
		Append(1).
		Append(2).
		Append(3).
		Append(4).
		Append(5)

	if !data.IsSorted(list.Asc[int]()) {
		t.Errorf("Expected list to be sorted")
	}
}

func Test_List_IsSorted_Fail(t *testing.T) {
	data := list.New[int]().
		Append(1).
		Append(2).
		Append(6).
		Append(4).
		Append(5)

	if data.IsSorted(list.Asc[int]()) {
		t.Errorf("Expected list to be sorted")
	}
}

func Test_List_IsSorted_Desc(t *testing.T) {
	data := list.New[int]().
		Append(5).
		Append(4).
		Append(3).
		Append(2).
		Append(1)

	if !data.IsSorted(list.Desc[int]()) {
		t.Errorf("Expected list to be sorted")
	}
}

func Test_List_IsSorted_Desc_Fail(t *testing.T) {
	data := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(2).
		Append(1)

	if data.IsSorted(list.Desc[int]()) {
		t.Errorf("Expected list to be sorted")
	}
}

func Test_List_Filter(t *testing.T) {
	data := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(7).
		Append(2).
		Append(1).
		Filter(list.Not(7))

	if data.Len() != 4 {
		t.Errorf("Expected 4, got %d", data.Len())
	}
}

func Test_List_Filter_Inclussive(t *testing.T) {
	data := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(7).
		Append(2).
		Append(1).
		Filter(list.Is(7))

	if data.Len() != 2 {
		t.Errorf("Expected 2, got %d", data.Len())
	}
}

func Test_List_Contains(t *testing.T) {
	found := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(7).
		Append(2).
		Append(1).
		Contains(list.Is(2))

	if !found {
		t.Errorf("Expected to find 2")
	}
}

func Test_List_Contains_Fail(t *testing.T) {
	found := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(7).
		Append(2).
		Append(1).
		Contains(list.Is(10))

	if found {
		t.Errorf("Expected to not find 10")
	}
}

func Test_Find(t *testing.T) {
	ten, err := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(7).
		Append(2).
		Append(10).
		Find(list.Is(10))

	if err != nil {
		t.Errorf("Expected to find 10")
	}

	if ten != 10 {
		t.Errorf("Expected 10, got %d", ten)
	}
}

func Test_Find_Fail(t *testing.T) {
	_, err := list.New[int]().
		Append(5).
		Append(4).
		Append(7).
		Append(7).
		Append(2).
		Append(10).
		Find(list.Is(11))

	if err == nil {
		t.Errorf("Expected to not find 11")
	}
}

func Test_Distinct(t *testing.T) {
	items := list.New[int]().
		Append(4).
		Append(7).
		Append(7).
		Append(7).
		Append(10).
		Append(10)

	expected := list.New[int]().
		Append(4).
		Append(7).
		Append(10)

	if !items.Distinct(list.Eq[int]()).Equals(expected, list.Eq[int]()) {
		t.Errorf("Expected %v, got %v", expected, items)
	}
}
