package data_structures

import "errors"

type Queue[T any] struct {
	list []T
}

type IQueue[T any] interface {
	Enqueue(i T)
	Dequeue() (firstElm T, err error)
}

func (q *Queue[T]) Enqueue(i T) {
	q.list = append(q.list, i)
}

func (q *Queue[T]) Len() int {
	return len(q.list)
}

func (q *Queue[T]) Dequeue() (firstElm T, err error) {
	if q.Len() == 0 {
		return firstElm, ErrDequeueEmptyQueue
	}

	firstElm = q.list[0]
	q.list = q.list[1:]

	return firstElm, nil
}

var ErrDequeueEmptyQueue error = errors.New("Empty Queue can not be Dequeued")
