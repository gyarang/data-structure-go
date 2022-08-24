package stack

import "testing"

func Test_Stack_Push_Length(t *testing.T) {
	stack := Stack[int]{}
	for i := 1; i < 10; i++ {
		stack.Push(i)
		if stack.length != i {
			t.Errorf("length of stack error, expect: %d, got: %d", i, stack.length)
		}
	}
}

func Test_Stack_Pop(t *testing.T) {
	stack := Stack[int]{}
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}

	for i := 10; i >= 0; i-- {
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
