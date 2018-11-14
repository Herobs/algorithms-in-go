package linkedlist

import (
	"fmt"
)

// Node linked list node
type Node struct {
	value interface{}
	next  *Node
}

// NewNode create new node with value and next link
func NewNode(value interface{}, next *Node) *Node {
	return &Node{value: value, next: next}
}

// String implement Stringer
func (node *Node) String() string {
	return fmt.Sprint(node.value)
}
