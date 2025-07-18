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
	newNode := &Node[K, V]{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}

	if m.size == 0 {
		m.root = newNode
		m.size++
	} else {
		lookup[K, V](key, nil, m.root, func(parent *Node[K, V], node *Node[K, V]) {
			diff := key.Compare(node.key)

			if diff < 0 {
				node.left = newNode
				m.size++
			} else if diff > 0 {
				node.right = newNode
				m.size++
			}
		})
	}
}

func (m *OrderedMap[K, V]) Erase(key K) {
	if m.size > 0 {
		lookup[K, V](key, nil, m.root, func(parent *Node[K, V], node *Node[K, V]) {
			if node != nil && key.Compare(node.key) == 0 {
				m.size--
				diff := key.Compare(parent.key)
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
		})
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
	m.root.traverse(action)
}

func lookup[K Comparable, V any](
	key K,
	parent *Node[K, V],
	node *Node[K, V],
	handler func(*Node[K, V], *Node[K, V]),
) {
	if node == nil {
		return
	}

	diff := key.Compare(node.key)

	if diff == 0 {
		handler(parent, node)
		return
	}

	if diff < 0 {
		if node.left != nil {
			lookup(key, node, node.left, handler)
		} else {
			handler(parent, node)
		}
	} else {
		if node.right != nil {
			lookup(key, node, node.right, handler)
		} else {
			handler(parent, node)
		}
	}
}

func (n *Node[K, V]) traverse(action func(K, V)) {
	if n == nil {
		return
	}

	n.left.traverse(action)
	action(n.key, n.value)
	n.right.traverse(action)
}

type Node[K, V any] struct {
	key         K
	value       V
	left, right *Node[K, V]
}

type Comparable interface {
	Compare(other Comparable) int
}
