package gobox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueuePut(t *testing.T) {
	queue := NewQueue()
	qLen := 10
	for i := 0; i < qLen; i++ {
		queue.Put(i)
	}
	assert.Equal(t, qLen, queue.container.Len(), "test Put() length")
	assert.Equal(t, 0, queue.container.Front().Value, "test append first element")
	assert.Equal(t, qLen-1, queue.container.Back().Value, "test append last element")
}

func TestQueueGet(t *testing.T) {
	queue := NewQueue()
	qLen := 10
	for i := 0; i < qLen; i++ {
		queue.Put(i)
	}
	assert.Equal(t, qLen, queue.container.Len(), "test Get() length before Put()")
	getValue := queue.Get()
	assert.Equal(t, qLen-1, queue.container.Len(), "test Get() length after Put()")
	assert.Equal(t, 0, getValue, "test Get() value")
}
