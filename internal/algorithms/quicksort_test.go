package algorithms

import (
"reflect"
"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{9, 3, 8, 1, 8, 5, 4}
	expected := []int{1, 3, 4, 5, 8, 8, 9}
	result := QuickSort(data)
	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Lists do not match. Expected: %v, Actual: %v", expected, result)
	}
}
