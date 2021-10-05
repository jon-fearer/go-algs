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
	expectedIndex := 6
	result, err := BinarySearch(sortedArray, searchVal)
	if err != nil {
		t.Errorf("Binary search error: %s", err.Error())
	}
	if result.resultValue != searchVal {
		t.Errorf("Binary search failed to locate search value. Expected: %d, Actual: %d", searchVal, result)
	}
	if result.resultIndex != expectedIndex {
		t.Errorf("Binary search failed to locate search value index. Expected: #{expectedIndex}, Actual: #{result.resultIndex}")
	}
}
