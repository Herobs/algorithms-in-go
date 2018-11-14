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

// String implement Stringer
func (node *Node) String() string {
	return fmt.Sprint(node.value)
}
