package stack

import "testing"

func Test_Stack_Push(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)

	for i := 2; i <= 10; i++ {
		if stack.last.value != i-1 {
			t.Errorf("last value error. expect: %d, got: %d", i-1, stack.last.value)
		}
		if stack.length != i-1 {
			t.Errorf("length error. expect: %d, got: %d", i-1, stack.length)
		}
		stack.Push(i)
	}
}

func Test_Stack_Pop(t *testing.T) {
	stack := Stack[int]{}
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}

	for i := 10; i >= 0; i-- {
		if stack.last.value != i {
			t.Errorf("pop last value error, expect: %d, got: %d", i, stack.last.value)
		}
		pop := stack.Pop()
		if i != pop {
			t.Errorf("pop stack error, expect: %d, got: %d", i, pop)
		}
	}

	pop := stack.Pop()
	if pop != 0 {
		t.Errorf("pop empty stack error, expect: 0, got: %d", pop)
	}
}

func Test_Stack_Length(t *testing.T) {
	stack := Stack[int]{}

	// push stack
	for i := 1; i <= 10; i++ {
		stack.Push(i)
		if stack.Length() != i {
			t.Errorf("length of stack error, expect: %d, got: %d", i, stack.Length())
		}
	}

	// pop stack
	for i := 10; i >= 1; i-- {
		stack.Pop()
		if stack.Length() != i-1 {
			t.Errorf("length of stack error, expect: %d, got: %d", i, stack.Length())
		}
	}

	// empty stack
	stack.Pop()
	if stack.Length() != 0 {
		t.Errorf("length of empty stack error, expect: 0, got: %d", stack.Length())
	}
}

func Test_Stack_Peek(t *testing.T) {
	stack := Stack[int]{}
	if stack.Peek() != 0 {
		t.Errorf("Peek of empty stack error: expect 0, got: %d", stack.Peek())
	}

	stack.Push(10)
	if stack.Peek() != 10 {
		t.Errorf("Peek of stack error: expect 10, got: %d", stack.Peek())
	}
}

func Test_Stack_IsEmpty(t *testing.T) {
	stack := Stack[int]{}
	if !stack.IsEmpty() {
		t.Errorf("IsEmpty of empty stack error, expect: true, got: false")
	}

	stack.Push(10)
	if stack.IsEmpty() {
		t.Errorf("IsEmpty of Not empty stack error, expect: false, got true")
	}
}
