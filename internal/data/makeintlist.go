package data

func MakeIntList(data []interface{}) []int {
	intList := make([]int, len(data))
	for i := range data {
		intList[i] = int(data[i].(float64))
	}
	return intList
}
