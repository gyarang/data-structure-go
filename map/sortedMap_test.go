package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedMap(t *testing.T) {
	sm := SortedMap[string, int]{}

	sm.Add("a", 1)
	v, ok := sm.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	sm.Add("c", 3)
	v, ok = sm.Get("c")
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	sm.Add("b", 2)
	v, ok = sm.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	assert.Equal(t, "a", sm.Arr[0].Key)
	assert.Equal(t, "b", sm.Arr[1].Key)
	assert.Equal(t, "c", sm.Arr[2].Key)
}

func TestSortedMap_Add(t *testing.T) {
	sm := SortedMap[string, int]{}
	sm.Add("a", 1)
	sm.Add("c", 3)
	sm.Add("b", 2)

	// default order
	assert.Equal(t, "a", sm.Arr[0].Key)
	assert.Equal(t, 1, sm.Arr[0].Value)
	assert.Equal(t, "b", sm.Arr[1].Key)
	assert.Equal(t, 2, sm.Arr[1].Value)
	assert.Equal(t, "c", sm.Arr[2].Key)
	assert.Equal(t, 3, sm.Arr[2].Value)

	// overlapping key
	sm.Add("b", 22)
	assert.Equal(t, "b", sm.Arr[1].Key)
	assert.Equal(t, 22, sm.Arr[1].Value)
}

func TestSortedMap_Get(t *testing.T) {
	sm := SortedMap[string, int]{}
	sm.Add("a", 1)
	sm.Add("b", 2)
	sm.Add("c", 3)

	v, ok := sm.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = sm.Get("b")
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = sm.Get("c")
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = sm.Get("d")
	assert.False(t, ok)
}

func TestSortedMap_Remove(t *testing.T) {
	sm := SortedMap[string, int]{}
	sm.Add("a", 1)

	v, ok := sm.Get("a")
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	ok = sm.Remove("a")
	assert.True(t, ok)
	v, ok = sm.Get("a")
	assert.False(t, ok)

	ok = sm.Remove("a")
	assert.False(t, ok)
}