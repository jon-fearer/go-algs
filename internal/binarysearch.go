package internal

import (
	"fmt"
	"math"
)

func BinarySearch(searchVal int) {
	parsed := ParseData("sorted-array-small.json")
	sortedArray := MakeIntList(parsed.([]interface{}))
	l := 0
	r := len(sortedArray) - 1
	for l <= r {
		m := math.Floor((float64(l) + float64(r))/2)
		fmt.Println(sortedArray[int(m)])
		if sortedArray[int(m)] < searchVal {
			l = int(m) + 1
		} else if sortedArray[int(m)] > searchVal {
			r = int(m) - 1
		} else {
			fmt.Println("found it!")
			return
		}
	}
}
