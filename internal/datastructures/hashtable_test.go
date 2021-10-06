package datastructures

import (
	"testing"
)

func TestHashTableEmptyGet(t *testing.T) {
	hashTable := NewHashTable()
	key := "something"
	getEmpty, err := hashTable.Get(key)
	if err != nil {
		t.Fatalf("Unable to 'Get' from hash table. Error: %s", err.Error())
	}
	if getEmpty != 0 {
		t.Errorf("Unexpected value received. Expected: %d, Actual: %d", 0, getEmpty)
	}
}

func TestHashTableValues(t *testing.T) {
	hashTable := NewHashTable()
	testKeyValues := []struct {
		key   string
		value int
	}{
		{"hello", 5},
		{"world", 6},
		{"hello", 7},
		{"goodbye", 5},
	}
	for i := 0; i < len(testKeyValues); i++ {
		key := testKeyValues[i].key
		value := testKeyValues[i].value
		err := hashTable.Set(key, value)
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
	}
}
