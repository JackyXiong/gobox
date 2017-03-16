package goboxes

import (
	"container/heap"
	"sync"
)

type item struct {
	priority int
	value    interface{}
}

type PriorityQueue struct {
	len   int
	items []*item
	mutex sync.Mutex
}

func NewPriorityQueue() *PriorityQueue {
	pq := new(PriorityQueue)
	pq.len = 0
	return pq
}

func (pq *PriorityQueue) Len() int {
	return pq.len
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].priority > pq.items[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(element interface{}) {
	newelement := element.(*item)
	pq.items = append(pq.items, newelement)
	pq.len++
}

func (pq *PriorityQueue) Pop() interface{} {
	item := pq.items[pq.len-1]
	pq.items = pq.items[:pq.len-1]
	return item
}

// Put a element into priority queue and set priority
func (pq *PriorityQueue) Put(vlaue interface{}, priority int) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	element := new(item)
	element.value = vlaue
	element.priority = priority
	heap.Push(pq, element)
}

// get a element from priority queue by priority
func (pq *PriorityQueue) Get() (vlaue interface{}) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	element := heap.Pop(pq)
	popElement := element.(*item)
	return popElement.value
}
