package gobox

import (
	"container/list"
	"sync"
)

// An unrestricted double-ended queue
type Deque struct {
	mutex     sync.Mutex
	container list.List
	len       int
}

func NewDeque() *Deque {
	return new(Deque)
}

func (q *Deque) Append(element interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.container.PushBack(element)
	q.len++
}

func (q *Deque) AppendLeft(element interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.container.PushFront(element)
	q.len++
}

func (q *Deque) First() interface{} {
	if q.Count() == 0 {
		return nil
	}
	return q.container.Front().Value
}

func (q *Deque) Last() interface{} {
	if q.Count() == 0 {
		return nil
	}
	return q.container.Back().Value
}

func (q *Deque) Extend(extendQ *Deque) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for e := extendQ.container.Front(); e != nil; e = e.Next() {
		q.container.PushBack(e.Value)
	}
	q.len = q.len + extendQ.len
}

func (q *Deque) ExtendLeft(extendQ *Deque) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for e := extendQ.container.Back(); e != nil; e = e.Prev() {
		q.container.PushFront(e.Value)
	}
	q.len = q.len + extendQ.len
}

func (q *Deque) Pop() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	back := q.container.Back()
	if back != nil {
		q.container.Remove(back)
		q.len--
		return back.Value
	}
	return nil
}

func (q *Deque) Count() int {
	return q.len
}
