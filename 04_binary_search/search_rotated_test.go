package binarysearch

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{3, 1}, 3, 0},
		{[]int{5, 1, 3}, 5, 0},
	}
	for _, tt := range tests {
		if got := SearchRotated(tt.nums, tt.target); got != tt.want {
			t.Errorf("SearchRotated(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
