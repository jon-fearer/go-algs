package datastructures

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	hashTable := NewHashTable()
	key := "hello"
	value := 5
	getEmpty, err := hashTable.Get(key)
	if err != nil {
		t.Fatalf("Unable to 'Get' from hash table. Error: %s", err.Error())
	}
	if getEmpty != 0 {
		t.Errorf("Unexpected value received. Expected: %d, Actual: %d", 0, getEmpty)
	}
	err = hashTable.Set(key, value)
	if err != nil {
		t.Fatalf("Unable to 'Set' on hash table. Error: %s", err.Error())
	}
	getValue, err := hashTable.Get(key)
	if err != nil {
		t.Fatalf("Unable to 'Get' from hash table. Error: %s", err.Error())
	}
	if getValue != value {
		t.Errorf("Unexpected value received. Expected: %d, Actual: %d", value, getValue)
	}
	key2 := "world"
	value2 := 6
	err = hashTable.Set(key2, value2)
	if err != nil {
		t.Fatalf("Unable to 'Set' on hash table. Error: %s", err.Error())
	}
	getValue, err = hashTable.Get(key2)
	if err != nil {
		t.Fatalf("Unable to 'Get' from hash table. Error: %s", err.Error())
	}
	if getValue != value2 {
		t.Errorf("Unexpected value received. Expected: %d, Actual: %d", value2, getValue)
	}
}
