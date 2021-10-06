package datastructures

type Node struct {
	data int
	next *Node
}

func IterateLinkedList(node *Node) ([]int, error) {
	currentNode := node
	data := []int{currentNode.data}
	for currentNode.next != nil {
		currentNode = currentNode.next
		data = append(data, currentNode.data)
	}
	return data, nil
}
