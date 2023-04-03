package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSparseSet_Add(t *testing.T) {
	ss := NewSparseSet[string, int]()

	ss.Add("a", 1)
	ss.Add("b", 2)

	i, ok := ss.sparse["a"]
	assert.True(t, ok)
	assert.Equal(t, 0, i)
	assert.Equal(t, ss.dense[0].Value, 1)
	i, ok = ss.sparse["b"]
	assert.True(t, ok)
	assert.Equal(t, 1, i)
	assert.Equal(t, ss.dense[1].Value, 2)

	ss.Add("a", 11)
	i, ok = ss.sparse["a"]
	assert.True(t, ok)
	assert.Equal(t, 0, i)
	assert.Equal(t, ss.dense[0].Value, 11)
}

func TestSparseSet_Get(t *testing.T) {
	ss := NewSparseSet[string, int]()

	ss.Add("a", 1)
	ss.Add("b", 2)

	v, ok := ss.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = ss.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = ss.Get("c")
	assert.False(t, ok)
}

func TestSparseSet_Remove(t *testing.T) {
	ss := NewSparseSet[string, int]()

	ss.Add("a", 1)
	ss.Add("b", 2)
	ss.Add("c", 3)

	ok := ss.Remove("b")
	assert.True(t, ok)

	_, ok = ss.Get("b")
	assert.False(t, ok)
}
