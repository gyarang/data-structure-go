package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap_Add(t *testing.T) {
	m := NewHashMap[int](4)
	m.Set("one", 1)
	h := m.hash("one")

	assert.Equal(t, 1, m.arr[h][0].value)

	m.Set("one", 11)
	assert.Equal(t, 11, m.arr[h][0].value)
}

func TestHashMap_Get(t *testing.T) {
	m := NewHashMap[int](4)

	item, ok := m.Get("one")
	assert.False(t, ok)
	assert.Equal(t, *new(int), item)

	m.Set("one", 1)
	item, ok = m.Get("one")
	assert.True(t, ok)
	assert.Equal(t, 1, item)
}

func TestHashMap_Delete(t *testing.T) {
	m := NewHashMap[int](4)
	m.Set("one", 1)

	m.Delete("one")
	item, ok := m.Get("one")
	assert.False(t, ok)
	assert.Equal(t, *new(int), item)
}
