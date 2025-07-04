package hw2

// go test ./hw2

type CircularQueue struct {
	values []int
	head   int
	tail   int
	length int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		head:   0,
		tail:   -1,
		length: 0,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	q.tail = (q.tail + 1) % len(q.values)
	q.values[q.tail] = value
	q.length++

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.head = (q.head + 1) % len(q.values)
	q.length--

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.head]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	v := q.values[q.tail]
	return v
}

func (q *CircularQueue) Empty() bool {
	return q.length == 0
}

func (q *CircularQueue) Full() bool {
	return q.length == len(q.values)
}
