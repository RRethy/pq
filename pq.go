package pq

// Evaluable TODO
type Evaluable interface {
	value() int
}

// MinPriorityQueue TODO
type MinPriorityQueue struct {
	members []Evaluable
}

// Push TODO
func (pq *MinPriorityQueue) Push(e Evaluable) {
	pq.members = append(pq.members, e)
	pq.fixUp()
}

// Pop TODO
func (pq *MinPriorityQueue) Pop() Evaluable {
	if pq.length() == 0 {
		return nil
	}
	retval := pq.members[pq.length()-1]
	pq.members = pq.members[:pq.length()-1]
	if pq.length() > 0 {
		pq.fixDown(0)
	}
	return retval
}

// Min TODO
func (pq *MinPriorityQueue) Min() Evaluable {
	if pq.length() == 0 {
		return nil
	}
	return pq.members[0]
}

func (pq *MinPriorityQueue) fixUp(i int) Evaluable {

}

func (pq *MinPriorityQueue) fixDown(i int) Evaluable {

}

func (pq *MinPriorityQueue) root() Evaluable {
	if pq.length() == 0 {
		return nil
	}
	return pq.members[0]
}

func (pq *MinPriorityQueue) length() int {
	return len(pq.members)
}
