package structures

// LIFO Stack
// Last-In-First-Out

type LinkedStack[T comparable] struct {
	list List[T]
}

func (s *LinkedStack[T]) Push(i T) {
	s.list.Push(i)
}

func (s *LinkedStack[T]) Pop() (i T, err error) {
	if s.list.Len() == 0 {
		return i, ErrPopEmptyStack
	}
	return s.list.Pop()
}

func (s *LinkedStack[T]) Len() int {
	return s.list.Len()
}
