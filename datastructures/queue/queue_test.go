package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	queue := New()

	assert.True(t, queue.IsEmpty())
}

func TestQueue_IsEmpty(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	assert.False(t, queue.IsEmpty())
}

func TestQueue_Peek(t *testing.T) {
	queue := New()

	assert.Nil(t, queue.Peek())
	queue.Enqueue(1)
	assert.Equal(t, 1, queue.Peek())
}

func TestQueue_Enqueue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	assert.Equal(t, 2, queue.list.Len())
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	assert.Equal(t, 1, queue.Dequeue())
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 3, queue.Dequeue())
	assert.Nil(t, queue.Dequeue())
	assert.Nil(t, queue.Peek())
}
