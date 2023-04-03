package _map

import "golang.org/x/exp/constraints"

type SparseSet[TKey constraints.Ordered, TValue any] struct {
	dense  []Element[TKey, TValue]
	sparse map[TKey]int
}

func NewSparseSet[TKey constraints.Ordered, TValue any]() *SparseSet[TKey, TValue] {
	return &SparseSet[TKey, TValue]{
		dense:  make([]Element[TKey, TValue], 0),
		sparse: make(map[TKey]int),
	}
}

func (s *SparseSet[TKey, TValue]) Add(key TKey, value TValue) {
	if e, ok := s.sparse[key]; ok {
		s.dense[e].Value = value
		return
	}

	s.dense = append(s.dense, Element[TKey, TValue]{Key: key, Value: value})
	s.sparse[key] = len(s.dense) - 1
}

func (s *SparseSet[TKey, TValue]) Get(key TKey) (TValue, bool) {
	if e, ok := s.sparse[key]; ok {
		return s.dense[e].Value, true
	}

	return *new(TValue), false
}

func (s *SparseSet[TKey, TValue]) Remove(key TKey) bool {
	if e, ok := s.sparse[key]; ok {
		s.dense = append(s.dense[:e], s.dense[e+1:]...)
		delete(s.sparse, key)
		return true
	}

	return false
}
