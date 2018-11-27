package priorityqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PriorityQueue(t *testing.T) {
	pq := &PriorityQueue{&_PriorityQueue{}}

	pq.Push(5, 5)
	pq.Push(10, 10)
	pq.Push(3, 3)
	pq.Push(2, 2)
	pq.Push(8, 8)

	assert.Equal(t, 5, pq.Len())
	assert.Equal(t, 10, pq.Pop())
	assert.Equal(t, 8, pq.Pop())
	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 0, pq.Len())
	assert.Nil(t, pq.Pop())
}
