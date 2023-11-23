package quicksort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	actual := []int{10, 80, 30, 90, 40, 50, 70}
	quicksort(actual)
	expected := []int{10, 30, 40, 50, 70, 80, 90}
	if reflect.DeepEqual(actual, expected) == false {
		t.FailNow()
	}
}
