package datastructures

import (
	"hash/fnv"
)

type KeyValueNode struct {
	data struct {
		key   string
		value int
	}
	next *KeyValueNode
}

func hashString(input string) (uint32, error) {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return 0, err
	}
	return hash.Sum32(), nil
}

func computeIndex(key string) (int, error) {
	hash, err := hashString(key)
	if err != nil {
		return 0, err
	}
	return int(hash % uint32(100)), nil
}

type IHashTable interface {
	Get(key string) (int, error)
	Set(key string, value int) error
}

type HashTable struct {
	data [100][]*KeyValueNode
}

func NewHashTable() IHashTable {
	return &HashTable{}
}

func (h *HashTable) Get(key string) (int, error) {
	index, err := computeIndex(key)
	if err != nil {
		return 0, err
	}
	kvLinkedList := h.data[index]
	var value int
	if kvLinkedList != nil {
		value, err = searchLinkedList(kvLinkedList, key)
		if err != nil {
			return 0, err
		}
	}
	return value, nil
}

func searchLinkedList(linkedList []*KeyValueNode, key string) (int, error) {
	currentNode := linkedList[0]
	var value int
	for currentNode != nil {
		if currentNode.data.key == key {
			value = currentNode.data.value
			break
		}
		currentNode = currentNode.next
	}
	return value, nil
}

func (h *HashTable) Set(key string, value int) error {
	index, err := computeIndex(key)
	if err != nil {
		return err
	}
	var next *KeyValueNode
	if h.data[index] != nil {
		next = h.data[index][0]
	}
	node := &KeyValueNode{struct {
		key   string
		value int
	}{key, value}, next}
	h.data[index] = append([]*KeyValueNode{node}, h.data[index]...)
	return nil
}
