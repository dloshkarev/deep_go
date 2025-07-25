package hw2

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	const (
		queueSize    = 3
		defaultValue = -1
	)
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front(defaultValue))
	assert.Equal(t, -1, queue.Back(defaultValue))
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front(defaultValue))
	assert.Equal(t, 3, queue.Back(defaultValue))

	assert.True(t, queue.Pop())
	assert.Equal(t, 2, queue.Front(defaultValue))
	assert.Equal(t, 3, queue.Back(defaultValue))
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front(defaultValue))
	assert.Equal(t, 4, queue.Back(defaultValue))

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
