package data_structures

type ListNode[T comparable] struct {
	Value T
	Next  *ListNode[T]
}

type ListHead[T comparable] struct {
	Node *ListNode[T]
}

func (ln *ListHead[T]) Insert(value T) {
	ln.Node = &ListNode[T]{Value: value, Next: ln.Node}
}

func (ln *ListHead[T]) Len() int {
	count := 1
	cur := ln.Node
	for cur.Next != nil {
		count++
		cur = cur.Next
	}

	return count
}

func (ln *ListHead[T]) ToList() []T {
	list := []T{}
	cur := ln.Node
	for cur.Next != nil {
		list = append(list, cur.Value)
		cur = cur.Next
	}

	return list
}
