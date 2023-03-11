package data_structures_test

import (
	"testing"

	"github.com/separovichabel/algorithms-datastructure/data_structures"
)

func TestQueueEnqueue(t *testing.T) {
	q := data_structures.Queue[int]{}

	q.Enqueue(1)

	if q.Len() != 1 {
		t.Errorf("Len should be 1 but got %v", q.Len())
	}

	q.Enqueue(2)
	if q.Len() != 2 {
		t.Errorf("Len should be 2 but got %v", q.Len())
	}
}

func TestQueueDenqueue(t *testing.T) {
	q := data_structures.Queue[int]{}

	var dequeued int
	var ok bool
	if dequeued, ok = q.Dequeue(); ok {
		t.Errorf("Empty Dequeued shoul not ok. But got %v, %v", dequeued, ok)
	}

	q.Enqueue(1)
	q.Enqueue(10)

	if dequeued, ok = q.Dequeue(); !ok {
		t.Errorf("Non-empty Dequeued shoul ok. But got %v, %v", dequeued, ok)
	}

	if dequeued != 1 {
		t.Errorf("First Dequeued value shoul be 1 but got %v", dequeued)
	}

	dequeued, ok = q.Dequeue()
	if dequeued != 10 {
		t.Errorf("Second Len should return 10 but got %v", dequeued)
	}
}
