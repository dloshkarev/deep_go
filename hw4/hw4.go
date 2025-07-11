package hw4

type OrderedMap struct {
	root *Node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		root: nil,
		size: 0,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	var parent, node *Node = nil, m.root

	for node != nil {
		if node.key == key {
			return
		}

		parent = node
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	newNode := &Node{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}

	if parent == nil {
		m.root = newNode
	} else {
		if key < parent.key {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}

	m.size++
}

func (m *OrderedMap) Erase(key int) {
	var parent, node *Node = nil, m.root

	for node != nil {
		if node.key == key {
			break
		}

		parent = node
		if key < node.key {
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

	isLeft := key < parent.key
	if node.right == nil {
		if isLeft {
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
		if isLeft {
			parent.left = prev
		} else {
			parent.right = prev
		}
	}
}

func (m *OrderedMap) Contains(key int) bool {
	node := m.root

	for node != nil {
		if node.key == key {
			return true
		}

		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	traverse(m.root, action)
}

func traverse(node *Node, action func(int, int)) {
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

type Node struct {
	key         int
	value       int
	left, right *Node
}
