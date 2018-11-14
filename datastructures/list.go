package datastructures

// TraverseFunc list traverse function
type TraverseFunc = func(interface{})

// TraverseFunc Find/Delete callback
type CallbackFunc = func(interface{}) bool

// List operations
type List interface {
	Prepend(value interface{})
	Append(value interface{})
	Delete(value interface{}) Node
	DeleteAll(value interface{})
	DeleteHead() Node
	DeleteTail() Node
	Reverse()
	Find(value interface{}, callback ...CallbackFunc) Node
	Traverse(callback TraverseFunc)
	TraverseReverse(callback TraverseFunc)
	FromIntArray(values []int)
	ToArray() []Node
	Head() Node
	Tail() Node
	Len() (length int)
	String() string
}
