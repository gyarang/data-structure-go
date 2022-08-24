package stack

type Element[T any] struct {
	prev  *Element[T]
	value T
}

type Stack[T any] struct {
	last   *Element[T]
	length int
}

func (s *Stack[T]) Push(v T) {
	e := &Element[T]{
		prev:  s.last,
		value: v,
	}

	s.last = e
	s.length++
}

func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		var zero T
		return zero
	}

	last := s.last
	s.last = last.prev
	s.length--
	return last.value
}
