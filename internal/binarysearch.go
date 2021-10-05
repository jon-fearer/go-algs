package internal

import (
	"math"
)

type BinarySearchResult struct {
	resultValue int
	resultIndex int
}

func BinarySearch(sortedArray []int, searchVal int) (BinarySearchResult, error) {
	left := 0
	right := len(sortedArray) - 1
	resultValue := 0
	resultIndex := 0
	for left <= right {
		middle := math.Floor((float64(left) + float64(right)) / 2)
		if sortedArray[int(middle)] < searchVal {
			left = int(middle) + 1
		} else if sortedArray[int(middle)] > searchVal {
			right = int(middle) - 1
		} else {
			resultIndex = int(middle)
			resultValue = sortedArray[int(middle)]
			break
		}
	}
	return BinarySearchResult{resultValue, resultIndex}, nil
}
