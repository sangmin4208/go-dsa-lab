package queue

import "errors"

type Queue[T any] struct {
	elements []T
}

var ErrEmptyQueue = errors.New("empty queue")

func (q *Queue[T]) Enqueue(val T) {
	q.elements = append(q.elements, val)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, ErrEmptyQueue
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, ErrEmptyQueue
	}

	return q.elements[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}
