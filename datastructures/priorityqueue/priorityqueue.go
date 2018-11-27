package priorityqueue

import "container/heap"

// Item with priority
type Item struct {
	value interface{}
	priority int
}

type _PriorityQueue []*Item

// Len length of queue
func (pq _PriorityQueue) Len() int {
	return len(pq)
}

// Less returns relation within index i and index j
func (pq _PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

// Swap swaps value at index i and index j
func (pq _PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push element at end
func (pq *_PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

// Pop remove and return the end element
func (pq *_PriorityQueue) Pop() interface{} {
	item := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]

	return item
}

// PriorityQueue priority queue
type PriorityQueue struct {
	store *_PriorityQueue
}

// Len length of queue
func (pq *PriorityQueue) Len() int {
	return pq.store.Len()
}

// Push element with priority
func (pq *PriorityQueue) Push(x interface{}, priority int) {
	heap.Push(pq.store, &Item{x, priority})
}

// Pop remove and return element with highest priority
func (pq *PriorityQueue) Pop() interface{} {
	if pq.store.Len() <= 0 {
		return nil
	}

	return heap.Pop(pq.store).(*Item).value
}
