package goboxes

type Queue struct {
	*Deque
}

func NewQueue() *Queue {
	return &Queue{NewDeque()}
}

func (q *Queue) Put(element interface{}) {
	q.Append(element)
}

func (q *Queue) Get() interface{} {
	return q.PopLeft()
}
