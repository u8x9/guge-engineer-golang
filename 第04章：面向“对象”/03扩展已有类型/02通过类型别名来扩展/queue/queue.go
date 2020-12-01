package queue

import (
	"errors"
)

type Queue []int

func (q *Queue) IsEmpty() bool {
	return q == nil || *q == nil || len(*q) == 0
}
func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	header := (*q)[0]
	*q = (*q)[1:]
	return header, nil
}

func (q *Queue) UnShift() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	lastIdx := len(*q) - 1
	tail := (*q)[lastIdx]
	*q = (*q)[:lastIdx]
	return tail, nil
}
