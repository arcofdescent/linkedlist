package queue

import "testing"

//import "fmt"

func TestCreate(t *testing.T) {
	CreateNew()
}

func TestED(t *testing.T) {
	q := CreateNew()

	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)

	if q.Dequeue() != 6 {
		t.Error(`Dequeue error`)
	}
	if q.Dequeue() != 7 {
		t.Error(`Dequeue error`)
	}
	if q.Dequeue() != 8 {
		t.Error(`Dequeue error`)
	}

	q.Enqueue(1)
	if q.Dequeue() != 1 {
		t.Error(`Dequeue error`)
	}
}
