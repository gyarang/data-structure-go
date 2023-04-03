package singlelinkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	cnt  int
	root *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) PushBack(v T) {
	n := &Node[T]{Value: v}

	if l.cnt == 0 {
		l.root = n
		l.tail = n
	} else {
		l.tail.Next = n
		l.tail = n
	}
	l.cnt++
}

func (l *LinkedList[T]) PushFront(v T) {
	n := &Node[T]{Value: v}

	if l.cnt == 0 {
		l.root = n
		l.tail = n
	} else {
		n.Next = l.root
		l.root = n
	}
	l.cnt++
}

func (l *LinkedList[T]) Front() *Node[T] { return l.root }

func (l *LinkedList[T]) Back() *Node[T] { return l.tail }

func (l *LinkedList[T]) Len() int { return l.cnt }

func (l *LinkedList[T]) GetAt(at int) *Node[T] {
	if at >= l.cnt {
		return nil
	}

	n := l.root
	for i := 0; i < at; i++ {
		n = n.Next
	}
	return n
}

func (l *LinkedList[T]) InsertAfter(n *Node[T], v T) {
	if !l.Contains(n) {
		return
	}

	newNode := &Node[T]{Value: v}
	newNode.Next, n.Next = n.Next, newNode
	l.cnt++
}

func (l *LinkedList[T]) Contains(n *Node[T]) bool {
	if n == nil {
		return false
	}

	for i := l.root; i != nil; i = i.Next {
		if i == n {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) Remove(n *Node[T]) bool {
	if n == nil {
		return false
	}

	if n == l.root {
		l.root = n.Next
		l.cnt--
		return true
	}

	for node := l.root; node != nil; node = node.Next {
		if node.Next == n {
			node.Next = n.Next
			l.cnt--
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) Array() []T {
	arr := make([]T, l.cnt)
	for i, n := 0, l.root; n != nil; i, n = i+1, n.Next {
		arr[i] = n.Value
	}
	return arr
}
