package doublylinkedlist

import (
	"fmt"
	"github.com/herobs/algorithms-in-go/datastructures"
	"strings"
)

// TraverseFunc list traverse function
type TraverseFunc = func(interface{})

// TraverseFunc Find/Delete callback
type CallbackFunc = func(interface{}) bool

// List doubly linked list
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
	node := NewNode(value, nil, list.head)

	if list.head == nil {
		list.tail = node
	} else {
		list.head.prev = node
	}
	list.head = node
}

// Append add value to list end
func (list *List) Append(value interface{}) {
	node := NewNode(value, list.tail, nil)

	if list.tail == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
}

// Delete first value from list
func (list *List) Delete(value interface{}) datastructures.Node {
	if list.head == nil {
		return nil
	}

	for curr := list.head; curr != nil; curr = curr.next {
		if curr.value == value {
			if curr == list.head {
				list.head = curr.next
				if list.head != nil {
					list.head.prev = nil
				} else {
					list.tail = nil
				}
			} else if curr == list.tail {
				list.tail = curr.prev
				list.tail.next = nil
			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
			}

			return curr
		}
	}

	return nil
}

// DeleteAll delete all from list
func (list *List) DeleteAll(value interface{}) {
	for {
		if deleted := list.Delete(value); deleted == nil {
			break
		}
	}
}

// DeleteHead delete head node from list
func (list *List) DeleteHead() datastructures.Node {
	if list.head == nil {
		return nil
	}

	deleted := list.head
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}

	return deleted
}

// DeleteTail delete tail node from list
func (list *List) DeleteTail() datastructures.Node {
	if list.tail == nil {
		return nil
	}

	deleted := list.tail
	list.tail = list.tail.prev
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}

	return deleted
}

// Reverse list in place
func (list *List) Reverse() {
	var prev, next *Node
	curr := list.head

	for curr != nil {
		next = curr.next

		curr.next = prev
		curr.prev = next

		prev = curr
		curr = next
	}

	list.tail = list.head
	list.head = prev
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

// TraverseReverse traverse list in reverse direction
func (list *List) TraverseReverse(callback TraverseFunc) {
	for curr := list.tail; curr != nil; curr = curr.prev {
		callback(curr)
	}
}

// FromIntArray append values to list
// note array of interface is not possible in go
func (list *List) FromIntArray(values []int) {
	for _, value := range values {
		list.Append(value)
	}
}

// ToArray get array of nodes form list
func (list *List) ToArray() []datastructures.Node {
	nodes := make([]datastructures.Node, 0)
	for curr := list.head; curr != nil; curr = curr.next {
		nodes = append(nodes, curr)
	}

	return nodes
}

// Head get head node
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
