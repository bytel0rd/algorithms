package main

type QueEntry struct {
	Data interface{}
	Next *QueEntry
}

func NewQueEntry(data interface{}) QueEntry {
	return QueEntry{
		Data: data,
		Next: nil,
	}
}

type Que struct {

	// the first item on the que, first removed
	first *QueEntry

	// last item on the que, which is last removed
	last *QueEntry

	// the length of the que
	length int
}

func NewQue() Que {
	return Que{first: nil, last: nil, length: 0}
}

func (q *Que) GetLength() int {
	return q.length
}

// returns the next person on the que
func (q *Que) Peek() *QueEntry {
	return q.first
}

// removes the first person on the Que same thing as POP
func (q *Que) Remove() bool {

	if q.length == 0 {
		return false
	}

	if q.length == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.Next
	}

	q.length--

	return true

}

func (q *Que) Add(entry QueEntry) {

	if q.length == 0 {
		q.first = &entry
		q.last = &entry
	} else {
		q.last.Next = &entry
		q.last = &entry
	}

	q.length++

}
