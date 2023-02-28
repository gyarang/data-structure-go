package _map

import (
	"hash/crc32"
)

type mapItem[T any] struct {
	key   string
	value T
}

type HashMap[T any] struct {
	size int
	arr  [][]mapItem[T]
}

func NewHashMap[T any](size int) HashMap[T] {
	return HashMap[T]{
		size: size,
		arr:  make([][]mapItem[T], size),
	}
}

func (m *HashMap[T]) Set(key string, value T) {
	h := m.hash(key)
	for i, item := range m.arr[h] {
		if item.key == key {
			m.arr[h][i].value = value
			return
		}
	}
	item := mapItem[T]{
		key:   key,
		value: value,
	}
	m.arr[h] = append(m.arr[h], item)
}

func (m *HashMap[T]) Get(key string) (T, bool) {
	h := m.hash(key)
	for _, item := range m.arr[h] {
		if item.key == key {
			return item.value, true
		}
	}
	return *new(T), false
}

func (m *HashMap[T]) Delete(key string) {
	h := m.hash(key)
	for i, item := range m.arr[h] {
		if item.key == key {
			m.arr[h] = append(m.arr[h][:i], m.arr[h][i+1:]...)
			return
		}
	}
}

func (m *HashMap[T]) hash(key string) int {
	index := crc32.ChecksumIEEE([]byte(key))
	return int(index) % m.size
}
