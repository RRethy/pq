package pq

import (
	"testing"
)

type E struct {
	val int
}

func (e *E) Value() int {
	return e.val
}

func TestPQ(t *testing.T) {
	input := []*E{
		&E{10},
		&E{5},
		&E{2},
		&E{8},
		&E{9},
		&E{7},
		&E{6},
		&E{1},
		&E{3},
		&E{4},
	}
	pq := new(PQ)
	for _, e := range input {
		pq.Push(e)
	}
	prev := -1
	var poppedOrder []*E
	for pq.Len() > 0 {
		popped := pq.Pop()
		poppedOrder = append(poppedOrder, popped.(*E))
		if popped.Value() < prev {
			t.Errorf("pq: %v | popped: %d | prev: %d\n", pq, popped, prev)
			break
		}
		prev = popped.Value()
	}

	for _, e1 := range poppedOrder {
		found := false
		for _, e2 := range input {
			if e1 == e2 {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Couldn't find match for *E with value %d\n", e1.Value())
		}
	}
}
