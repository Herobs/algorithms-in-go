package datastructures

// Stack operations
type Stack interface {
	IsEmpty() bool
	Peek() interface{}
	Push(value interface{})
	Pop() interface{}
	String() string
}
