package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	stack := New()

	assert.True(t, stack.IsEmpty())
}

func TestStack_IsEmpty(t *testing.T) {
	stack := New()

	stack.Push(1)
	assert.False(t, stack.IsEmpty())
}

func TestStack_Peek(t *testing.T) {
	stack := New()

	assert.Nil(t, stack.Peek())
	stack.Push(1)
	assert.Equal(t, 1, stack.Peek())
}

func TestStack_Enstack(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.list.Len())
}

func TestStack_Destack(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Nil(t, stack.Pop())
	assert.Nil(t, stack.Peek())
}
