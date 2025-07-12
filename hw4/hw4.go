package hw4

type OrderedMap[K Comparable, V any] struct {
	root *Node[K, V]
	size int
}

func NewOrderedMap[K Comparable, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		root: nil,
		size: 0,
	}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	var parent, node *Node[K, V] = nil, m.root
	var diff = 0

	for node != nil {
		diff = key.Compare(node.key)

		if diff == 0 {
			return
		}

		parent = node
		if diff < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	newNode := &Node[K, V]{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}

	if parent == nil {
		m.root = newNode
	} else {
		if diff < 0 {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}

	m.size++
}

func (m *OrderedMap[K, V]) Erase(key K) {
	var parent, node *Node[K, V] = nil, m.root
	var diff = 0

	for node != nil {
		diff = key.Compare(node.key)

		if diff == 0 {
			break
		}

		parent = node
		if diff < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	if node == nil {
		return
	}

	m.size--

	if parent == nil {
		m.root = nil
		return
	}

	diff = key.Compare(parent.key)
	if node.right == nil {
		if diff < 0 {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
	} else {
		prev, minNode := node.right, node.right.left
		for minNode != nil {
			prev = minNode
			minNode = minNode.left
		}

		prev.left = nil
		if diff < 0 {
			parent.left = prev
		} else {
			parent.right = prev
		}
	}
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	node := m.root
	var diff = 0

	for node != nil {
		diff = key.Compare(node.key)

		if diff == 0 {
			return true
		}

		if diff < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

func (m *OrderedMap[K, V]) Size() int {
	return m.size
}

func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	traverse(m.root, action)
}

func traverse[K Comparable, V any](node *Node[K, V], action func(K, V)) {
	if node == nil {
		return
	}

	if node.left != nil {
		traverse(node.left, action)
	}

	action(node.key, node.value)

	if node.right != nil {
		traverse(node.right, action)
	}
}

type Node[K, V any] struct {
	key         K
	value       V
	left, right *Node[K, V]
}

type Comparable interface {
	Compare(other Comparable) int
}
