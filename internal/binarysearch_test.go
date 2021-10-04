package internal

import (
	"github.com/jon-fearer/go-algs/internal/data"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	parsed, err := data.ParseData("sorted-array-small.json")
	if err != nil {
		t.Fatalf("Unable to parse test data. Error: %s", err.Error())
	}
	sortedArray := data.MakeIntList(parsed.([]interface{}))
	searchVal := 7
	result, err := BinarySearch(sortedArray, searchVal)
	if err != nil {
		t.Errorf("Binary search error: %s", err.Error())
	}
	if result != searchVal {
		t.Errorf("Binary search failed to locate search value. Expected: %d, Actual: %d", searchVal, result)
	}
}
