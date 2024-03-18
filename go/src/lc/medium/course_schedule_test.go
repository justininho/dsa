package medium

import (
	"testing"
)

func TestCourseSchedule(t *testing.T) {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
	}

	for _, tt := range tests {
		got := canFinish(tt.numCourses, tt.prerequisites)
		if got != tt.want {
			t.Errorf("canFinish(%v, %v) return %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
		}
	}
}
