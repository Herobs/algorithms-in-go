package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_String(t *testing.T) {
	node := &Node{value: 10}

	assert.Equal(t, fmt.Sprint(node), "10")
}

func TestList_Append(t *testing.T) {
	list := &List{}

	list.Append(1)
	assert.Equal(t, list.head, list.tail)
	assert.Len(t, list.ToArray(), 1)
	assert.Equal(t, "1", fmt.Sprint(list))

	list.Append(2)
	assert.Len(t, list.ToArray(), 2)
	assert.Equal(t, "1, 2", fmt.Sprint(list))
}

func TestList_Prepend(t *testing.T) {
	list := &List{}

	list.Prepend(1)
	assert.Equal(t, list.head, list.tail)
	assert.Len(t, list.ToArray(), 1)
	assert.Equal(t, "1", fmt.Sprint(list))

	list.Prepend(2)
	assert.Len(t, list.ToArray(), 2)
	assert.Equal(t, "2, 1", fmt.Sprint(list))
}

func TestList_Delete(t *testing.T) {
	list := &List{}

	list.Append(1)
	list.Append(2)
	list.Prepend(2)
	assert.Len(t, list.ToArray(), 3)
	assert.Equal(t, "2, 1, 2", fmt.Sprint(list))

	list.Delete(2)
	assert.Equal(t, "1, 2", fmt.Sprint(list))
	list.Delete(2)
	assert.Equal(t, "1", fmt.Sprint(list))
}

func TestList_DeleteAll(t *testing.T) {
	list := &List{}

	list.Append(1)
	list.Append(2)
	list.Prepend(2)
	list.DeleteAll(2)
	assert.Equal(t, "1", fmt.Sprint(list))
}

func TestList_DeleteHead(t *testing.T) {
	list := &List{}

	list.Append(1)
	list.DeleteHead()
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	list.Append(2)
	list.Append(3)
	list.DeleteHead()
	assert.Equal(t, list.head, list.tail)
	assert.Equal(t, "3", fmt.Sprint(list))

	list.Append(4)
	list.Append(5)
	list.DeleteHead()
	assert.Equal(t, "4, 5", fmt.Sprint(list))
}

func TestList_DeleteTail(t *testing.T) {
	list := &List{}

	list.Append(1)
	list.DeleteTail()
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	list.Append(2)
	list.Append(3)
	list.DeleteTail()
	assert.Equal(t, list.head, list.tail)
	assert.Equal(t, "2", fmt.Sprint(list))

	list.Append(4)
	list.Append(5)
	list.DeleteTail()
	assert.Equal(t, "2, 4", fmt.Sprint(list))
}

func TestList_FromIntArray(t *testing.T) {
	list := &List{}

	list.FromIntArray([]int{1, 2, 3})
	assert.Equal(t, "1, 2, 3", fmt.Sprint(list))
}

func TestList_ToArray(t *testing.T) {
	list := &List{}

	list.FromIntArray([]int{1, 2, 3})
	nodes := list.ToArray()
	assert.Len(t, nodes, 3)

}

func TestList_String(t *testing.T) {
	list := &List{}

	list.FromIntArray([]int{1, 2, 3})
	assert.Equal(t, "1, 2, 3", list.String())
}

func TestList_Reverse(t *testing.T) {
	list := &List{}

	list.Reverse()
	assert.Equal(t, "", list.String())

	list.Append(1)
	list.Reverse()
	assert.Equal(t, "1", list.String())

	list.Append(2)
	list.Reverse()
	assert.Equal(t, "2, 1", list.String())

	list.Append(3)
	list.Reverse()
	assert.Equal(t, "3, 1, 2", list.String())
}

func TestList_ReverseRecursive(t *testing.T) {
	list := &List{}

	list.ReverseRecursive()
	assert.Equal(t, "", list.String())

	list.Append(1)
	list.ReverseRecursive()
	assert.Equal(t, "1", list.String())

	list.Append(2)
	list.ReverseRecursive()
	assert.Equal(t, "2, 1", list.String())

	list.Append(3)
	list.ReverseRecursive()
	assert.Equal(t, "3, 1, 2", list.String())
}

func TestList_Find(t *testing.T) {
	list := &List{}

	list.FromIntArray([]int{1, 2, 3})
	assert.Nil(t, list.Find(4))
	assert.Equal(t, 1, list.Find(1).Value())
}

func TestList_Traverse(t *testing.T) {
	list := &List{}

	list.FromIntArray([]int{1, 2, 3})

	values := make([]int, 0)
	list.Traverse(func(value interface{}) {
		values = append(values, value.(*Node).value.(int))
	})
	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestList_TraverseReverse(t *testing.T) {
	list := &List{}

	list.FromIntArray([]int{1, 2, 3})

	values := make([]int, 0)
	list.TraverseReverse(func(value interface{}) {
		values = append(values, value.(*Node).value.(int))
	})
	assert.Equal(t, []int{3, 2, 1}, values)
}
