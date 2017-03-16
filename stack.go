package goboxes

type Stack struct {
	*Deque
}

func NewStack() *Stack {
	return &Stack{NewDeque()}
}

func (q *Stack) Put(element interface{}) {
	q.Append(element)
}

func (q *Stack) Get() interface{} {
	return q.Pop()
}
