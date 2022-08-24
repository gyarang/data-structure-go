package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		set: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(v T) {
	s.set[v] = struct{}{}
}

func (s *Set[T]) Delete(v T) {
	delete(s.set, v)
}

func (s *Set[T]) Contain(v T) bool {
	_, c := s.set[v]
	return c
}

func (s *Set[T]) Length() int {
	return len(s.set)
}

func (s *Set[T]) Items() []T {
	items := make([]T, len(s.set))
	i := 0
	for k := range s.set {
		items[i] = k
		i++
	}

	return items
}
