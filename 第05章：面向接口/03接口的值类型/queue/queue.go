package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) UnShift() interface{} {
	lastIndex := len(*q) - 1
	v := (*q)[lastIndex]
	*q = (*q)[:lastIndex]
	return v
}
