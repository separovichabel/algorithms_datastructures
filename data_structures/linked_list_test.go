package data_structures_test

import (
	"testing"

	"github.com/separovichabel/algorithms-datastructure/data_structures"
)

func TestListNodeInsert(t *testing.T) {
	head := data_structures.ListHead[int]{nil}

	head.Insert(0)

	if head.Node.Value != 0 {
		t.Errorf("Value should be 0 but got %v", head.Node.Value)
	}
}

type lenCase[T comparable] struct {
	arr      []T
	expected int
}

var lenCases []lenCase[int] = []lenCase[int]{
	{
		[]int{1, 2, 3, 4},
		4,
	},
	{
		[]int{1},
		1,
	},
	{
		[]int{0, 1, 2},
		3,
	},
}

func TestListNodeLen(t *testing.T) {
	for _, tcase := range lenCases {
		tListNode := produceData(tcase.arr)
		resp := tListNode.Len()

		if resp != tcase.expected {
			t.Errorf("Value should be %v but got %v", tcase.expected, resp)
		}
	}
}

type toListCase[T comparable] struct {
	arr      []T
	expected []T
}

var toListCases []toListCase[int] = []toListCase[int]{
	{
		[]int{1, 2, 3, 4},
		[]int{4, 3, 2, 1},
	},
	{
		[]int{1},
		[]int{1},
	},
	{
		[]int{0, 1, 2},
		[]int{2, 1, 0},
	},
}

func TestListNodeToList(t *testing.T) {
	for _, tcase := range toListCases {
		tListNode := produceData(tcase.arr)
		resp := tListNode.ToList()

		for i, item := range resp {
			if tcase.expected[i] != item {
				t.Errorf("Resp[%v] should be %v but got %v", i, tcase.expected, item)
			}
		}
	}
}

func produceData[T comparable](data []T) data_structures.ListHead[T] {

	head := data_structures.ListHead[T]{Node: nil}

	for _, d := range data {
		head.Insert(d)
	}

	return head
}
