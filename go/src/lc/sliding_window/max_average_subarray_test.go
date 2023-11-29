package lc

import "testing"

func TestFindMaxAverage(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{1, 12, -5, -6, 50, 3}, 1, 50},
		{[]int{1, 12, -5, -6, 50, 3}, 2, 26.5},
		{[]int{1, 12, -5, -6, 50, 3}, 3, 16.333333333333332},
		{[]int{1, 12, -5, -6, 50, 3}, 5, 10.8},
		{[]int{1, 12, -5, -6, 50, 3}, 6, 8.333333333333334},
		{[]int{1, 12, -5, -6, 50, 3}, 7, 7.285714285714286},
		{[]int{1, 12, -5, -6, 50, 3}, 8, 6.5},
	}

	for _, tt := range tests {
		got := findMaxAverage(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("findMaxAverage(%v, %v) return %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
