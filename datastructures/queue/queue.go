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

// Peek get the bottom value without removing
func (queue Queue) Peek() interface{} {
	if node := queue.list.Head(); node != nil {
		return node.Value()
	}

	return nil
}

// Enqueue add value on the top
func (queue *Queue) Enqueue(value interface{}) {
	queue.list.Append(value)
}

// Dequeue remove the bottom value and return
func (queue *Queue) Dequeue() interface{} {
	if node := queue.list.DeleteHead(); node != nil {
		return node.Value()
	}

	return nil
}

// String implement Stringer
func (queue *Queue) String() string {
	return queue.list.String()
}
