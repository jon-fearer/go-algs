package datastructures

import (
	"reflect"
	"testing"
)

func TestIterateLinkedList(t *testing.T) {
	node3Data := 578
	node3 := Node{node3Data, nil}
	node2Data := 8457
	node2 := Node{node2Data, &node3}
	node1Data := 9
	node1 := Node{node1Data, &node2}
	expectedList := []int{node1Data, node2Data, node3Data}
	actualList, err := IterateLinkedList(&node1)
	if err != nil {
		t.Fatalf("Unable to iterate linked list. Error: %s", err.Error())
	}
	if reflect.DeepEqual(actualList, expectedList) == false {
		t.Errorf("Lists do not match. Expected: %v, Actual: %v", expectedList, actualList)
	}
}
