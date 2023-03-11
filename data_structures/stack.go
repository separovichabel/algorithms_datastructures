package data_structures

import "errors"

// LIFO Stack
// Last-In-First-Out

type Stack[T any] struct {
	list []T
}

// type IStack[T any] interface {
// 	Push(i T)
// 	Pop() (i T, err error)
//  Len() int
// }

func (s *Stack[T]) Push(i T) {
	s.list = append(s.list, i)
}

func (s *Stack[T]) Pop() (i T, err error) {
	if len(s.list) == 0 {
		return i, ErrPopEmptyStack
	}
	i = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return
}

func (s *Stack[T]) Len() int {
	return len(s.list)
}

var ErrPopEmptyStack error = errors.New("Empty Stack Pop Error")
