package queue

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(content interface{}) {
	*q = append(*q, content)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	element := (*q)[0]
	if len(*q) > 1 {
		*q = (*q)[1:]
	}
	return element
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	
	return (*q)[0]
}