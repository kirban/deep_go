package main

import (
	"cmp"
)

/**
*  Идея упорядоченного словаря заключается в том, что он будет реализован
*  на основе бинарного дерева поиска (BST). Дерево будет строиться только по ключам
*  элементов, значения элементов при построении дерева не учитываются.
*  Элементы с одинаковыми ключами в упорядоченном словаре хранить нельзя
 */

type Node[K cmp.Ordered, V any] struct {
	left, right *Node[K, V]
	key         K
	value       V
}

type OrderedMap[K cmp.Ordered, V any] struct {
	root *Node[K, V]
}

func NewNode[K cmp.Ordered, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		left:  nil,
		right: nil,
		key:   key,
		value: value,
	}
}

// Min поиск ноды с минимальным ключом в поддереве
func (node *Node[K, V]) Min() *Node[K, V] {
	if node == nil {
		return nil
	}

	n := node
	for n.left != nil {
		n = node.left
	}

	return n
}

// NewOrderedMap создать упорядоченный словарь
func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		root: nil,
	}
}

// Insert добавить элемент в словарь
func (m *OrderedMap[K, V]) Insert(key K, value V) {
	var targetNode *Node[K, V] = nil

	for node := m.root; node != nil; {
		if key == node.key {
			node.value = value
			return
		} else {
			targetNode = node
			if key < node.key {
				node = node.left
			} else {
				node = node.right
			}
		}
	}

	if targetNode == nil {
		m.root = NewNode(key, value)
	} else {
		if key < targetNode.key {
			targetNode.left = NewNode(key, value)
		} else {
			targetNode.right = NewNode(key, value)
		}
	}
}

// Erase удалить элемент из словаря
func (m *OrderedMap[K, V]) Erase(key K) bool {
	if m.root == nil {
		return false
	}

	var targetNode, parentNode *Node[K, V] = m.root, nil
	isLeft := false

	for targetNode != nil && targetNode.key != key {
		parentNode = targetNode
		if key < targetNode.key {
			targetNode = targetNode.left
			isLeft = true
		} else if key > targetNode.key {
			targetNode = targetNode.right
			isLeft = false
		}
	}

	if targetNode == nil {
		// node not found
		return false
	}

	if targetNode.left == nil && targetNode.right == nil {
		if targetNode == m.root {
			m.root = nil
		} else if isLeft {
			parentNode.left = nil
		} else {
			parentNode.right = nil
		}
	} else if targetNode.right == nil {
		if targetNode == m.root {
			m.root = targetNode.left
		} else if isLeft {
			parentNode.left = targetNode.left
		} else {
			parentNode.right = targetNode.left
		}
	} else if targetNode.left == nil {
		if targetNode == m.root {
			m.root = targetNode.right
		} else if isLeft {
			parentNode.left = targetNode.right
		} else {
			parentNode.right = targetNode.right
		}
	} else {
		minNode := targetNode.right.Min()

		m.Erase(minNode.key)

		targetNode.key = minNode.key
		targetNode.value = minNode.value
	}

	return true
}

// Contains проверить существование элемента в словаре
func (m *OrderedMap[K, V]) Contains(key K) bool {
	for node := m.root; node != nil; {
		if node.key == key {
			return true
		} else if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		}
	}

	return false
}

// Size получить количество элементов в словаре
func (m *OrderedMap[K, V]) Size() int {
	count := 0

	if m.root == nil {
		return count
	}

	m.ForEach(func(_ K, _ V) {
		count++
	})

	return count
}

// ForEach применить функцию к каждому элементу словаря от меньшего к большему
func (m *OrderedMap[K, V]) ForEach(f func(K, V)) {
	var stack []*Node[K, V]
	current := m.root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		f(current.key, current.value)

		current = current.right
	}
}

func main() {}
