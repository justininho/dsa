package lc

import (
	"slices"
	"testing"
)

func TestCourseSchedule2(t *testing.T) {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{2, [][]int{{1, 0}, {0, 1}}, []int{}},
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, []int{2, 1, 0}},
		{3, [][]int{{1, 0}, {1, 2}, {0, 1}}, []int{}},
	}

	for _, tt := range tests {
		got := findOrder(tt.numCourses, tt.prerequisites)
		if !slices.Equal(got, tt.want) {
			t.Errorf("findOrder(%v, %v) return %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
		}
	}
}
