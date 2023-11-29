package lc

import "testing"

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		//{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		//{[]int{2, 2, 2, 3, 1, 1, 4, 1}, 4, 2},
	}
	for _, tt := range tests {
		if got := maxOperations(tt.nums, tt.k); got != tt.want {
			t.Errorf("maxOperations(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
