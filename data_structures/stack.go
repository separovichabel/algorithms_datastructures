package data_structures

// LIFO Stack
// Last-In-First-Out

type Stack[T any] struct {
	list []T
}

func (s *Stack[T]) Push(i T) {
	s.list = append(s.list, i)
}

func (s *Stack[T]) Pop() (i T) {
	i = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return
}

func (s *Stack[T]) Len() int {
	return len(s.list)
}
