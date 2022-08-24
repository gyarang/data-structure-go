package set

import (
	"sort"
	"testing"
)

func Test_Set_Add(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"default", []int{1, 2, 3}, []int{1, 2, 3}},
		{"redundant", []int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet[int]()
			for _, v := range tt.input {
				set.Add(v)
			}
			items := set.Items()
			if !compareIntSlices(items, tt.expect) {
				t.Errorf("test case %s failed, expect: %v, get: %v", tt.name, tt.expect, items)
			}
		})
	}
}

func compareIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
