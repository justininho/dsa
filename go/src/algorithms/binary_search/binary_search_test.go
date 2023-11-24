package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{10, 30, 40, 50, 70, 80, 90}
	if search(nums, 30) != 1 {
		t.FailNow()
	}
	if search(nums, 90) != 6 {
		t.FailNow()
	}
	if search(nums, 100) != -1 {
		t.FailNow()
	}
}
