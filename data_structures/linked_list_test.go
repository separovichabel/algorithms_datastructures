package data_structures_test

import (
	"testing"

	"github.com/separovichabel/algorithms-datastructure/data_structures"
)

func TestListNodeInsert(t *testing.T) {
	head := data_structures.List[int]{nil}

	head.Insert(0)

	if head.Node.Value != 0 {
		t.Errorf("Value should be 0 but got %v", head.Node.Value)
	}
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

type lenCase struct {
	arr      []int
	expected int
}

var lenCases []lenCase = []lenCase{
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

type toListCase struct {
	arr      []int
	expected []int
}

var toListCases []toListCase = []toListCase{
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

func produceData[T comparable](data []T) data_structures.List[T] {

	head := data_structures.List[T]{Node: nil}

	for _, d := range data {
		head.Insert(d)
	}

	return head
}
