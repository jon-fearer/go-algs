package internal

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

func ParseData(filename string) interface{} {
	filepath := path.Join("internal", "data", filename)
	dat, err := ioutil.ReadFile(filepath)
	Check(err)
	var parsed interface{}
	err = json.Unmarshal(dat, &parsed)
	Check(err)
	return parsed
}

func MakeIntList(dat []interface{}) []int {
	intList := make([]int, len(dat))
	for i := range dat {
		intList[i] = int(dat[i].(float64))
	}
	return intList
}
