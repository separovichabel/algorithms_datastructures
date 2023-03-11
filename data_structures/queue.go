package data_structures

type Queue[T any] struct {
	list []T
}

func (q *Queue[T]) Enqueue(i T) {
	q.list = append(q.list, i)
}

func (q *Queue[T]) Len() int {
	return len(q.list)
}

func (q *Queue[T]) Dequeue() (firstElm T, ok bool) {
	if q.Len() == 0 {
		return firstElm, false
	}

	firstElm = q.list[0]
	q.list = q.list[1:]

	return firstElm, true
}
