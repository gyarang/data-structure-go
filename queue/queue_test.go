package queue

import "testing"

func Test_Queue_Add(t *testing.T) {
	q := &Queue[int]{}
	for i := 1; i < 10; i++ {
		q.Add(i)
		if !CheckQueue(q, 1, i, i) {
			t.Errorf("queue add failed, expect - head: %d, tail: %d, length: %d, got: %d, %d, %d", 1, i, i, q.head.value, q.tail.value, q.length)
		}
	}
}

func Test_Queue_Poll(t *testing.T) {
	q := &Queue[int]{}
	for i := 1; i <= 10; i++ {
		q.Add(i)
	}

	for i := 1; i <= 10; i++ {
		v := q.Poll()
		if v != i {
			t.Errorf("queue poll failed, expect: %d, got: %d", i, v)
		}
	}

	v := q.Poll()
	if v != 0 {
		t.Errorf("poll empty pueue failed, expect: 0, got: %d", v)
	}
}

func Test_Queue_Peek(t *testing.T) {
	q := &Queue[int]{}

	for i := 1; i <= 10; i++ {
		q.Add(i)
	}

	for i := 1; i <= 10; i++ {
		v := q.Peek()
		if v != i {
			t.Errorf("queue peek failed, expect: %d, got: %d", i, v)
		}
		q.Poll()
	}

	v := q.Peek()
	if v != 0 {
		t.Errorf("poll empty peek failed, expect: 0, got: %d", v)
	}
}

func Test_Queue_IsEmpty(t *testing.T) {
	q := &Queue[int]{}
	if !q.IsEmpty() {
		t.Errorf("queue IsEmpty failed, expect: true, got: false")
	}
	q.Add(1)
	if q.IsEmpty() {
		t.Errorf("queue IsEmpty failed, expect: false, got: true")
	}
}

func Test_Queue_Add_Cleared_Queue(t *testing.T) {
	q := &Queue[int]{}
	q.Add(1)
	q.Poll()
	q.Add(100)

	if !CheckQueue(q, 100, 100, 1) {
		t.Errorf("queue add failed, expect - head: %d, tail: %d, length: %d, got: %d, %d, %d", 100, 100, 1, q.head.value, q.tail.value, q.length)
	}
}

func CheckQueue(q *Queue[int], head, tail, length int) bool {
	if q.head.value != head || q.tail.value != tail || q.length != length {
		return false
	}
	return true
}
