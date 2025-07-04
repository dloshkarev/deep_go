package hw2

// go test ./hw2

type CircularQueue[T any] struct {
	values []T
	head   int
	tail   int
	length int
}

func NewCircularQueue[T any](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
		head:   0,
		tail:   -1,
		length: 0,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	q.tail = (q.tail + 1) % len(q.values)
	q.values[q.tail] = value
	q.length++

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	q.head = (q.head + 1) % len(q.values)
	q.length--

	return true
}

func (q *CircularQueue[T]) Front(defaultValue T) T {
	if q.Empty() {
		return defaultValue
	}

	return q.values[q.head]
}

func (q *CircularQueue[T]) Back(defaultValue T) T {
	if q.Empty() {
		return defaultValue
	}

	v := q.values[q.tail]
	return v
}

func (q *CircularQueue[T]) Empty() bool {
	return q.length == 0
}

func (q *CircularQueue[T]) Full() bool {
	return q.length == len(q.values)
}
