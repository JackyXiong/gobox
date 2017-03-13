package gobox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	deque := NewDeque()
	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	assert.Equal(t, qLen, deque.container.Len(), "test Append length")
	assert.Equal(t, 0, deque.container.Front().Value, "test append first element")
	assert.Equal(t, qLen-1, deque.container.Back().Value, "test append last element")
}

func TestAppendLeft(t *testing.T) {
	deque := NewDeque()
	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.AppendLeft(i)
	}
	assert.Equal(t, qLen, deque.container.Len(), "test AppendLeft length")
	assert.Equal(t, qLen-1, deque.container.Front().Value, "test AppendLeft first element")
	assert.Equal(t, 0, deque.container.Back().Value, "test AppendLeft last element")
}

func TestFirst(t *testing.T) {
	deque := NewDeque()
	assert.Equal(t, nil, deque.Last(), "test method First() empty deque")

	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	assert.Equal(t, 0, deque.First(), "test method First()")
}

func TestLast(t *testing.T) {
	deque := NewDeque()
	assert.Equal(t, nil, deque.Last(), "test method Last() empty deque")

	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	assert.Equal(t, qLen-1, deque.Last(), "test method Last()")
}

func TestExtend(t *testing.T) {
	aDeque := NewDeque()
	bDeque := NewDeque()

	qLen := 10
	for i := 0; i < qLen; i++ {
		aDeque.Append(i)
		bDeque.Append(i + 100)
	}
	aDeque.Extend(bDeque)
	assert.Equal(t, qLen*2, aDeque.Count(), "test Extend()")
	assert.Equal(t, 0, aDeque.First(), "test Extend()")
	assert.Equal(t, 109, aDeque.Last(), "test Extend()")
}

func TestExtendLeft(t *testing.T) {
	aDeque := NewDeque()
	bDeque := NewDeque()

	qLen := 10
	for i := 0; i < qLen; i++ {
		aDeque.Append(i)
		bDeque.Append(i + 100)
	}
	aDeque.ExtendLeft(bDeque)
	assert.Equal(t, qLen*2, aDeque.Count(), "test Extend()")
	assert.Equal(t, 100, aDeque.First(), "test ExtendLeft()")
	assert.Equal(t, 9, aDeque.Last(), "test ExtendLeft()")
}

func TestPop(t *testing.T) {
	deque := NewDeque()
	assert.Equal(t, nil, deque.Pop(), "test method Pop() empty deque")

	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	assert.Equal(t, qLen-1, deque.Pop(), "test method Pop()")
}

func TestPopLeft(t *testing.T) {
	deque := NewDeque()
	assert.Equal(t, nil, deque.PopLeft(), "test method PopLeft() empty deque")

	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	assert.Equal(t, 0, deque.PopLeft(), "test method PopLeft()")
}

func TestCount(t *testing.T) {
	deque := NewDeque()
	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	assert.Equal(t, qLen, deque.Count(), "test method Count()")
}

func TestRemove(t *testing.T) {
	deque := NewDeque()
	qLen := 10
	for i := 0; i < qLen; i++ {
		deque.Append(i)
	}
	element := 3
	removed_element := deque.Remove(element)
	assert.Equal(t, qLen-1, deque.Count(), "test method Remove()")
	assert.Equal(t, removed_element, element, "test method Remove()")
}
