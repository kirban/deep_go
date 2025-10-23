package main

import (
	"unsafe"
)

type CircularQueue struct {
	values []int
	front  *int
	rear   *int
}

// NewCircularQueue - создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	cq := CircularQueue{
		values: make([]int, size),
		front:  nil,
		rear:   nil,
	}

	return cq
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	// if is first elem
	if q.front == nil && q.rear == nil {
		q.front = &q.values[0]
		q.rear = &q.values[0]
	} else {
		// increase rear pointer by one
		if q.rear == &q.values[len(q.values)-1] {
			q.rear = &q.values[0]
		} else {
			q.rear = (*int)(unsafe.Add(unsafe.Pointer(q.rear), 1<<3))
		}
	}

	*q.rear = value

	return true
}

// Pop - удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	*q.front = 0

	if q.front != q.rear {
		if q.front == &q.values[len(q.values)-1] {
			q.front = &q.values[0]
		} else {
			q.front = (*int)(unsafe.Add(unsafe.Pointer(q.front), 1<<3))
		}
	} else {
		q.front, q.rear = nil, nil
	}
	return true
}

// Front - получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return *q.front
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	return *q.rear
}

// Empty - проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.front == nil && q.rear == nil
}

// Full - проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.front == &q.values[0] &&
		q.rear == &q.values[len(q.values)-1]
}

func main() {

	cq := NewCircularQueue(3)
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	cq.Pop()
	cq.Push(4)
}
