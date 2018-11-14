package queue

import (
	"github.com/herobs/algorithms-in-go/datastructures"
	"github.com/herobs/algorithms-in-go/datastructures/linkedlist"
)

// Queue fifo queue
type Queue struct {
	list datastructures.List
}

// New create new queue
func New() *Queue {
	return &Queue{
		list: linkedlist.New(),
	}
}

// IsEmpty check empty
func (queue *Queue) IsEmpty() bool {
	return queue.list.Head() == nil
}

// Peek get the top value without removing
func (queue Queue) Peek() interface{} {
	return queue.list.Head()
}

// Enqueue add value on the top
func (queue *Queue) Enqueue(value interface{}) {
	queue.list.Append(value)
}

// Dequeue remove the top value and return
func (queue *Queue) Dequeue() interface{} {
	return queue.list.DeleteHead()
}

// String implement Stringer
func (queue *Queue) String() string {
	return queue.list.String()
}
