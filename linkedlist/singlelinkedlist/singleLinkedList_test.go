package singlelinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test for SingleLinkedList
func TestSingleLinkedList(t *testing.T) {
	l := &LinkedList[int]{}
	l.PushBack(2)
	l.PushBack(3)

	assert.Equal(t, []int{2, 3}, l.Array())

	l.PushFront(1)
	assert.Equal(t, []int{1, 2, 3}, l.Array())

	n := l.GetAt(0)
	assert.Equal(t, 1, n.Value)

	l.InsertAfter(n, 11)
	assert.Equal(t, []int{1, 11, 2, 3}, l.Array())

	contains := l.Contains(n)
	assert.True(t, contains)

	ok := l.Remove(n)
	assert.True(t, ok)
	assert.Equal(t, []int{11, 2, 3}, l.Array())

	length := l.Len()
	assert.Equal(t, 3, length)

}
