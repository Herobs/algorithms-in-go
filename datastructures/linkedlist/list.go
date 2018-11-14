package linkedlist

import (
	"fmt"
	"strings"

	"github.com/herobs/algorithms-in-go/datastructures"
)

// TraverseFunc traverse list node
type TraverseFunc = func(interface{})

// TraverseFunc Find/Delete callback
type CallbackFunc = func(interface{}) bool

// List linked list
type List struct {
	head *Node
	tail *Node
}

var _ datastructures.List = New()

// New create new list
func New() *List {
	return &List{}
}

// Prepend add value to list front
func (list *List) Prepend(value interface{}) {
	node := NewNode(value, list.head)
	list.head = node

	if list.tail == nil {
		list.tail = node
	}
}

// Append add value to list end
func (list *List) Append(value interface{}) {
	node := NewNode(value, nil)

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	list.tail = node
}

// Delete first value from list
func (list *List) Delete(value interface{}) (deleted datastructures.Node) {
	if list.head == nil {
		return nil
	}

	if list.head.value == value {
		deleted = list.head
		list.head = list.head.next
		return
	}

	curr := list.head
	if curr != nil {
		for ; curr.next != nil; curr = curr.next {
			if curr.next.value == value {
				deleted = curr.next
				curr.next = curr.next.next
				break
			}
		}
	}

	if list.tail == deleted {
		list.tail = curr
	}

	return
}

// DeleteAll delete all from list
func (list *List) DeleteAll(value interface{}) {
	for {
		if deleted := list.Delete(value); deleted == nil {
			return
		}
	}
}

// DeleteHead delete head node from list
func (list *List) DeleteHead() (deleted datastructures.Node) {
	if list.head == nil {
		return nil
	}

	deleted = list.head
	list.head = list.head.next

	if list.tail == deleted {
		list.tail = nil
	}

	return
}

// DeleteTail delete tail node from list
func (list *List) DeleteTail() (deleted datastructures.Node) {
	if list.tail == nil {
		return nil
	}

	deleted = list.tail
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		return
	}

	curr := list.head
	for ; curr.next != nil; curr = curr.next {
		if curr.next.next == nil {
			curr.next = nil
			break
		}
	}
	list.tail = curr

	return
}

// Reverse list in place
func (list *List) Reverse() {
	var prev, next *Node
	curr := list.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	list.tail = list.head
	list.head = prev
}

// reverseRecursive recursive function for reversing
func (list *List) reverseRecursive(head *Node) *Node {
	if head.next == nil {
		return head
	}

	tail := list.reverseRecursive(head.next)
	head.next = nil
	tail.next = head

	return head
}

// ReverseRecursive list in place recursively
func (list *List) ReverseRecursive() {
	if list.head == nil {
		return
	}

	tail := list.reverseRecursive(list.head)
	list.head = list.tail
	list.tail = tail
}

// Find first value in list
func (list *List) Find(value interface{}, callback ...CallbackFunc) datastructures.Node {
	if list.head == nil {
		return nil
	}

	for curr := list.head; curr != nil; curr = curr.next {
		if len(callback) > 0 {
			if callback[0](curr.value) {
				return curr
			}
		} else {
			if curr.value == value {
				return curr
			}
		}
	}

	return nil
}

// Traverse visit all nodes in list
func (list *List) Traverse(callback TraverseFunc) {
	for curr := list.head; curr != nil; curr = curr.next {
		callback(curr)
	}
}

// traverseReverse recursive function for traversing
func (list *List) traverseReverse(head *Node, callback TraverseFunc) {
	if head == nil {
		return
	}

	list.traverseReverse(head.next, callback)
	callback(head)
}

// TraverseReverse traverse list in reverse direction
func (list *List) TraverseReverse(callback TraverseFunc) {
	list.traverseReverse(list.head, callback)
}

// FromIntArray append values to list
// note array of interface is not possible in go
func (list *List) FromIntArray(values []int) {
	for _, value := range values {
		list.Append(value)
	}
}

// ToArray get array of values form list
func (list *List) ToArray() []datastructures.Node {
	nodes := make([]datastructures.Node, 0)
	for curr := list.head; curr != nil; curr = curr.next {
		nodes = append(nodes, curr)
	}

	return nodes
}

// Head get head value
func (list *List) Head() datastructures.Node {
	if list.head == nil {
		return nil
	}

	return list.head
}

// Tail get tail value
func (list *List) Tail() datastructures.Node {
	if list.tail == nil {
		return nil
	}

	return list.tail
}

// Len get the length of list
func (list *List) Len() (length int) {
	for curr := list.head; curr != nil; curr = curr.next {
		length++
	}

	return
}

// String implement Stringer
func (list *List) String() string {
	nodes := list.ToArray()
	values := make([]string, 0, len(nodes))
	for _, node := range nodes {
		values = append(values, fmt.Sprint(node))
	}

	return strings.Join(values, ", ")
}
