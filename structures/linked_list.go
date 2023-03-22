package structures

import "errors"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Node  *Node[T]
	lengh int
}

func (ln *List[T]) Push(value T) {
	ln.Node = &Node[T]{Value: value, Next: ln.Node}
	ln.lengh++
}

func (ln *List[T]) Pop() (value T, err error) {
	if ln.Node == nil {
		return value, ErrPopEmptyList
	}
	value = ln.Node.Value
	ln.Node = ln.Node.Next
	ln.lengh--
	return
}

func (ln *List[T]) Len() int {
	return ln.lengh
}

func (ln *List[T]) ToList() []T {
	list := []T{}
	cur := ln.Node
	for cur.Next != nil {
		list = append(list, cur.Value)
		cur = cur.Next
	}

	return list
}

var ErrPopEmptyList error = errors.New("Empty List Pop Error")
