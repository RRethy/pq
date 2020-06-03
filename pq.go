package pq

import (
	"fmt"
	"math"
)

// Evaluable TODO
type Evaluable interface {
	Value() int
}

// PQ TODO
type PQ struct {
	members []Evaluable
}

func (pq *PQ) String() string {
	var mapped []int
	for _, e := range pq.members {
		mapped = append(mapped, e.Value())
	}
	return fmt.Sprintf("%v", mapped)
}

// Push TODO
func (pq *PQ) Push(e Evaluable) {
	pq.members = append(pq.members, e)
	pq.fixUp(pq.Len() - 1)
}

// Pop TODO
func (pq *PQ) Pop() Evaluable {
	if pq.Len() == 0 {
		return nil
	}
	pq.swap(0, pq.Len()-1)
	retval := pq.members[pq.Len()-1]
	pq.members = pq.members[:pq.Len()-1]
	if pq.Len() > 0 {
		pq.fixDown(0)
	}
	return retval
}

// Min TODO
func (pq *PQ) Min() Evaluable {
	if pq.Len() == 0 {
		return nil
	}
	return pq.members[0]
}

// Len TODO
func (pq *PQ) Len() int {
	return len(pq.members)
}

func (pq *PQ) fixUp(i int) {
	for pq.parent(i) != -1 && pq.valueAt(pq.parent(i)) > pq.valueAt(i) {
		pq.swap(pq.parent(i), i)
		i = pq.parent(i)
	}
}

func (pq *PQ) fixDown(i int) {
	for !pq.leaf(i) {
		minChildIndex := pq.indexMinChild(i)
		if pq.valueAt(i) < pq.valueAt(minChildIndex) {
			break
		}
		pq.swap(i, minChildIndex)
		i = minChildIndex
	}
}

func (pq *PQ) root() Evaluable {
	if pq.Len() == 0 {
		return nil
	}
	return pq.members[0]
}

func (pq *PQ) leaf(i int) bool {
	return 2*i+1 >= pq.Len()
}

// TODO return an error instead of -1 if i is a leaf
func (pq *PQ) indexMinChild(i int) int {
	if pq.Len() > 2*i+1 {
		if pq.Len() > 2*i+2 && pq.valueAt(2*i+1) > pq.valueAt(2*i+2) {
			return 2*i + 2
		}
		return 2*i + 1
	}
	return -1 // i is a leaf
}

func (pq *PQ) valueAt(i int) int {
	if i >= pq.Len() {
		return math.MaxInt32
	}
	return pq.members[i].Value()
}

func (pq *PQ) swap(i, j int) {
	pq.members[i], pq.members[j] = pq.members[j], pq.members[i]
}

func (pq *PQ) parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}
