package main

import (
	"fmt"
)

type CircularQueue struct {
	values []int
	front  int
	rear   int
	count  int
}

// NewCircularQueue - создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	cq := CircularQueue{
		values: make([]int, size),
		front:  -1,
		rear:   -1,
	}

	return cq
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	if q.rear == -1 {
		q.front = 0
	}

	q.rear = (q.rear + 1) % cap(q.values)
	q.values[q.rear] = value
	q.count += 1

	return true
}

// Pop - удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.values[q.front] = 0
	q.front = (q.front + 1) % cap(q.values)

	q.count -= 1

	return true
}

// Front - получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.front]
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.rear]
}

// Empty - проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

// Full - проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.count == cap(q.values)
}

func main() {
	q := NewCircularQueue(3)

	fmt.Println(q.Push(1))
}
