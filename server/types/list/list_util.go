package list

import (
	"fmt"
)

// @generate str
// @replace ListItem >> string
// @replace ListType >> StrList

type ListItem = interface{}

type ListType []ListItem

func (list ListType) IndexOf(item ListItem, isEq func(ListItem, ListItem) bool) (int, error) {
	for idx, it := range list.AsArray() {
		if isEq(it, item) {
			return idx, nil
		}
	}

	return -1, fmt.Errorf("object %v does not exist in list", item)
}

func (list ListType) Contains(item ListItem, isEq func(ListItem, ListItem) bool) bool {
	_, err := list.IndexOf(item, isEq)
	return err == nil
}

func (list ListType) Len() int {
	return len(list.AsArray())
}

func (list ListType) AsArray() []ListItem {
	return []ListItem(list)
}

func (list ListType) Filter(predicate func(int, ListItem) bool) ListType {
	out := make([]ListItem, 0)
	for idx, it := range list.AsArray() {
		if predicate(idx, it) {
			out = append(out, it)
		}
	}

	return out
}
