package datastructures

// Queue operations
type Queue interface {
	IsEmpty() bool
	Peek() interface{}
	Enqueue(value interface{})
	Dequeue() interface{}
	String() string
}
