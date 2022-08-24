package queue

type Element[T any] struct {
	next  *Element[T]
	value T
}

type Queue[T any] struct {
	head, tail *Element[T]
	length     int
}

func (q *Queue[T]) Add(v T) {
	e := &Element[T]{
		next:  nil,
		value: v,
	}

	if q.length == 0 {
		q.head = e
		q.tail = e
	} else {
		q.tail.next = e
		q.tail = e
	}

	q.length++
}

func (q *Queue[T]) Poll() T {
	if q.length == 0 {
		return ZeroValue[T]()
	}

	start := q.head
	q.head = start.next
	start.next = nil
	q.length--

	return start.value
}

func (q *Queue[T]) Peek() T {
	if q.length == 0 {
		return ZeroValue[T]()
	}
	return q.head.value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func ZeroValue[T any]() T {
	var zero T
	return zero
}
