package stack

import (
	"github.com/herobs/algorithms-in-go/datastructures"
	"github.com/herobs/algorithms-in-go/datastructures/linkedlist"
)

// Stack filo stack
type Stack struct {
	list datastructures.List
}

// New create new queue
func New() *Stack {
	return &Stack{
		list: linkedlist.New(),
	}
}

// IsEmpty check empty
func (stack *Stack) IsEmpty() bool {
	return stack.list.Head() == nil
}

// Peek get the top value without removing
func (stack *Stack) Peek() interface{} {
	if node := stack.list.Tail(); node != nil {
		return node.Value()
	}

	return nil
}

// Push add value on the top
func (stack *Stack) Push(value interface{}) {
	stack.list.Append(value)
}

// Pop remove the top value and return
func (stack *Stack) Pop() interface{} {
	if node := stack.list.DeleteTail(); node != nil {
		return node.Value()
	}
	return nil
}

// String implement Stringer
func (stack *Stack) String() string {
	return stack.list.String()
}
