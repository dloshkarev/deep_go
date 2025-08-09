package hw12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()

	_, err := scheduler.GetTask()
	assert.Error(t, err, ErrEmptyHeap)

	scheduler.ChangeTaskPriority(2, 500)

	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task, err := scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task5, task)

	task, err = scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)

	task, err = scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task1, task)

	task, err = scheduler.GetTask()
	assert.NoError(t, err)
	assert.Equal(t, task3, task)
}
