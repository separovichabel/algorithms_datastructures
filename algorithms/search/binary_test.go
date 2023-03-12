package search_test

import (
	"testing"

	"github.com/separovichabel/algorithms_datastructures/algorithms/search"
)

type Case struct {
	name     string
	input1   []int
	input2   int
	expected int
}

var cases []Case = []Case{
	{
		"1",
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		5,
		4,
	},
	{
		"2",
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		9,
		-1,
	},
	{
		"3",
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		1,
		0,
	},
	{
		"4",
		[]int{1, 2, 3, 4, 6, 7, 8},
		5,
		-1,
	},
}

func TestBinary(t *testing.T) {
	for _, c := range cases {
		output := search.Binary(c.input1, c.input2)
		if output != c.expected {
			t.Errorf("Case Name %v: (output) %v != %v (expected)", c.name, output, c.expected)
		}
	}
}
