package structures_test

import (
	"testing"

	"github.com/separovichabel/algorithms_datastructures/structures"
)

func TestQueueEnqueue(t *testing.T) {
	q := structures.Queue[int]{}

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
	q := structures.Queue[int]{}

	var dequeued int
	var err error

	q.Enqueue(1)
	q.Enqueue(10)

	if dequeued, err = q.Dequeue(); err != nil {
		t.Errorf("Non-empty Dequeued shoul not error. But got \"%v\"", err.Error())
	}

	if dequeued != 1 {
		t.Errorf("First Dequeued value shoul be 1 but got %v", dequeued)
	}

	dequeued, err = q.Dequeue()
	if dequeued != 10 {
		t.Errorf("Second Len should return 10 but got %v", dequeued)
	}
}

func TestQueueDenqueueEmpty(t *testing.T) {
	q := structures.Queue[int]{}

	var dequeued int
	var err error
	if dequeued, err = q.Dequeue(); err == nil {
		t.Errorf("Empty Dequeued shoul error. But got %v, %v", dequeued, err)
	}
}
