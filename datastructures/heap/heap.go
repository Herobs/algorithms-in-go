package heap

import (
	"sort"
)

// Interface same as official container/heap package
type Interface interface {
	sort.Interface
	Push(interface{})
	Pop() interface{}
}

// Init build heap structure
func Init(heap Interface) {
	n := heap.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(heap, i, n)
	}
}

// Push pushes element onto heap
func Push(heap Interface, x interface{}) {
	heap.Push(x)
	up(heap, heap.Len()-1)
}

// Pop remove the top element and return
func Pop(heap Interface) interface{} {
	n := heap.Len() - 1
	heap.Swap(0, n)
	down(heap, 0, n)
	return heap.Pop()
}

// Remove removes the element at index i from the heap
func Remove(heap Interface, i int) interface{} {
	n := heap.Len() - 1
	if n != i {
		heap.Swap(i, n)
		if !down(heap, i, n) {
			up(heap, i)
		}
	}

	return heap.Pop()
}

// Fix keep heap structure after the element at index i has changed its value
func Fix(heap Interface, i int) {
	if !down(heap, i, heap.Len()) {
		up(heap, i)
	}
}

func up(heap Interface, i int) {
	for {
		p := (i - 1) / 2
		if p == i || !heap.Less(i, p) {
			break
		}
		heap.Swap(i, p)
		i = p
	}
}

func down(heap Interface, i, n int) bool {
	p := i
	for {
		left := p*2 + 1
		if left >= n {
			break
		}
		pos := left
		if right := left + 1; right < n && heap.Less(right, left) {
			pos = right
		}
		if !heap.Less(pos, p) {
			break
		}
		heap.Swap(pos, p)
		p = pos
	}
	return p > i
}
