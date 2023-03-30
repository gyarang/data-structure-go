package _map

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type Element[TKey constraints.Ordered, TValue any] struct {
	Key   TKey
	Value TValue
}

type SortedMap[TKey constraints.Ordered, TValue any] struct {
	Arr []Element[TKey, TValue]
}

func (s *SortedMap[TKey, TValue]) Add(key TKey, value TValue) {
	index := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if index < len(s.Arr) && s.Arr[index].Key == key {
		s.Arr[index].Value = value
		return
	}

	s.Arr = append(s.Arr[:index], append([]Element[TKey, TValue]{{Key: key, Value: value}}, s.Arr[index:]...)...)
}

func (s *SortedMap[TKey, TValue]) Get(key TKey) (TValue, bool) {
	index := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if index < len(s.Arr) && s.Arr[index].Key == key {
		return s.Arr[index].Value, true
	}

	return *new(TValue), false
}

func (s *SortedMap[TKey, TValue]) Remove(key TKey) bool {
	index := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if index < len(s.Arr) && s.Arr[index].Key == key {
		s.Arr = append(s.Arr[:index], s.Arr[index+1:]...)
		return true
	}

	return false
}
