package data

func MakeIntList(dat []interface{}) []int {
	intList := make([]int, len(dat))
	for i := range dat {
		intList[i] = int(dat[i].(float64))
	}
	return intList
}
