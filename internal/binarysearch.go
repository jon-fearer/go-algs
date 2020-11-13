package internal

import "fmt"

func BinarySearch() {
	parsed := ParseData("sorted-array-small.json")
	sortedArray := MakeIntList(parsed.([]interface{}))
	fmt.Println(sortedArray)
}
