package hw12

import (
	"container/heap"
	"errors"
)

var (
	ErrEmptyHeap = errors.New("no tasks")
)

type Task struct {
	Identifier int
	Priority   int
}

type HTask struct {
	Task
	heapPriority int
	idx          int
}

type Scheduler struct {
	tasks    *TaskHeap
	taskById map[int]*HTask
}

func NewScheduler() Scheduler {
	tasks := TaskHeap([]*HTask{})
	heap.Init(&tasks)

	return Scheduler{
		tasks:    &tasks,
		taskById: make(map[int]*HTask),
	}
}

func (s *Scheduler) AddTask(task Task) {
	hTask := HTask{task, task.Priority, 0}
	s.taskById[task.Identifier] = &hTask
	heap.Push(s.tasks, &hTask)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	if t, exists := s.taskById[taskID]; exists {
		t.heapPriority = newPriority
		heap.Fix(s.tasks, t.idx)
	}
}

func (s *Scheduler) GetTask() (Task, error) {
	if s.tasks.Len() == 0 {
		return Task{}, ErrEmptyHeap
	}
	task := heap.Pop(s.tasks).(*HTask).Task
	delete(s.taskById, task.Identifier)
	return task, nil
}

type TaskHeap []*HTask

func (h TaskHeap) Len() int {
	return len(h)
}

func (h TaskHeap) Less(i, j int) bool {
	return h[i].heapPriority > h[j].heapPriority
}

func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}

func (h *TaskHeap) Push(x interface{}) {
	item := x.(*HTask)
	item.idx = len(*h)
	*h = append(*h, x.(*HTask))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
