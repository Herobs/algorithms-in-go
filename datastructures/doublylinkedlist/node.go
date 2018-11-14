package doublylinkedlist

import (
	"fmt"
)

// Node doubly linked list node
type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

// NewNode create new node with value, prev and next link
func NewNode(value interface{}, prev *Node, next *Node) *Node {
	return &Node{value: value, prev: prev, next: next}
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
