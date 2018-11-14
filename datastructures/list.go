package datastructures

// TraverseFunc list traverse function
type TraverseFunc = func(interface{})

// List operations
type List interface {
	Prepend(value interface{})
	Append(value interface{})
	Delete(value interface{}) interface{}
	DeleteAll(value interface{})
	DeleteHead() interface{}
	DeleteTail() interface{}
	Reverse()
	Find(value interface{}) interface{}
	Traverse(callback TraverseFunc)
	TraverseReverse(callback TraverseFunc)
	FromIntArray(values []int)
	ToArray() []interface{}
	Head() interface{}
	Tail() interface{}
	Len() (length int)
	String() string
}
