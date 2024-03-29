package structures_test

import (
	"testing"

	"github.com/separovichabel/algorithms_datastructures/structures"
)

func TestStackPush(t *testing.T) {
	s := structures.Stack[int]{}
	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Expected len %v but got %v", 2, s.Len())
	}
}

func TestStackPop(t *testing.T) {
	s := structures.Stack[int]{}
	s.Push(3)
	s.Push(4)
	resp, err := s.Pop()

	if err != nil {
		t.Errorf("Should not error when Poping a non-empty stack")
	}

	if resp != 4 {
		t.Errorf("Expected Pop to return %v but got %v", 4, resp)
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := structures.Stack[int]{}
	_, err := s.Pop()

	if err == nil {
		t.Errorf("Should error when Poping a empty Stack")
	}
}

func TestStackLen(t *testing.T) {
	s := structures.Stack[string]{}
	s.Push("foo")
	s.Push("bar")
	resp := s.Len()
	if resp != 2 {
		t.Errorf("Expected len %v but got %v", 2, resp)
	}
	s.Pop()
	resp = s.Len()
	if resp != 1 {
		t.Errorf("Expected len %v but got %v", 1, resp)
	}
}
