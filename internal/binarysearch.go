package internal

import (
	"math"
)

func BinarySearch(sortedArray []int, searchVal int) (int, error) {
	left := 0
	right := len(sortedArray) - 1
	result := 0
	for left <= right {
		middle := math.Floor((float64(left) + float64(right)) / 2)
		if sortedArray[int(middle)] < searchVal {
			left = int(middle) + 1
		} else if sortedArray[int(middle)] > searchVal {
			right = int(middle) - 1
		} else {
			result = sortedArray[int(middle)]
			break
		}
	}
	return result, nil
}
