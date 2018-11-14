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

// Value get node value
func (node *Node) Value() interface{} {
	return node.value
}

// SetValue set node value
func (node *Node) SetValue(value interface{}) {
	node.value = value
}

// String implement Stringer
func (node *Node) String() string {
	return fmt.Sprint(node.value)
}
