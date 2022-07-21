package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type UniqueQueue struct {
	queue       []interface{}
	front, rear int
	length      int
}

func New(size int) Queue {
	return &UniqueQueue{
		queue:  make([]interface{}, size),
		front:  0,
		length: 0,
		rear:   -1,
	}
}

func (q *UniqueQueue) isFull() bool {
	return q.length == len(q.queue)
}

func (q *UniqueQueue) isEmpty() bool {
	return q.length == 0
}

func (q *UniqueQueue) Push(key interface{}) {
	var isAvailable bool = !q.isFull()

	if isAvailable {
		q.length = q.length + 1

		q.rear = 1 + q.rear
		q.queue[q.rear] = key
	} else {
		q.Pop()

		q.length = q.length + 1

		q.rear = 1 + q.rear
		q.queue[q.rear] = key
	}
}

func (q *UniqueQueue) Pop() interface{} {
	if q.isEmpty() {
		return nil
	}

	var tmp interface{} = q.queue[q.front]

	for i := 0; i < q.length-1; i++ {
		q.queue[i] = q.queue[i+1]
	}

	q.length = q.length - 1

	q.queue[q.length] = nil
	q.rear--

	return tmp
}

func (q *UniqueQueue) Contains(key interface{}) bool {
	for _, val := range q.queue {
		if val == key {
			return true
		}
	}

	return false
}

func (q *UniqueQueue) Len() int {
	return q.length
}

func (q *UniqueQueue) Keys() []interface{} {
	var keys []interface{} = []interface{}{}

	for i := 0; i < q.length; i++ {
		keys = append(keys, q.queue[i])
	}

	return keys
}
