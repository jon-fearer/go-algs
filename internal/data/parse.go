package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

func ParseData(filename string) (interface{}, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filepath := path.Join(workingDir, "..", "data", filename)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
