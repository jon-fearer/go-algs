package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	dataSet := []struct{
		input, output []int
	}{
		{[]int{9, 3, 8, 1, 8, 5, 4}, []int{1, 3, 4, 5, 8, 8, 9}},
		{[]int{4, 5, 8, 1, 8, 3, 9}, []int{1, 3, 4, 5, 8, 8, 9}},
		{[]int{1, 3, 4, 5, 8, 8, 9}, []int{1, 3, 4, 5, 8, 8, 9}},
		{[]int{9, 8, 8, 5, 4, 3, 1}, []int{1, 3, 4, 5, 8, 8, 9}},
	}
	for i := 0; i < len(dataSet); i++ {
		result, err := MergeSort(dataSet[i].input)
		if err != nil {
			t.Errorf("Merge sort error: %s", err.Error())
		}
		if reflect.DeepEqual(result, dataSet[i].output) == false {
			t.Errorf("Lists do not match. Expected: %v, Actual: %v", dataSet[i].output, result)
		}
	}
}
