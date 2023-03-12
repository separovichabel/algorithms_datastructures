package structures

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Node *Node[T]
}

func (ln *List[T]) Insert(value T) {
	ln.Node = &Node[T]{Value: value, Next: ln.Node}
}

func (ln *List[T]) Len() int {
	count := 1
	cur := ln.Node
	for cur.Next != nil {
		count++
		cur = cur.Next
	}

	return count
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
