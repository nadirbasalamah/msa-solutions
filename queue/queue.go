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

// New returns a fixed size queue
func New(size int) Queue {
	return &UniqueQueue{
		queue:  make([]interface{}, size),
		front:  0,
		length: 0,
		rear:   -1,
	}
}

// isFull checks if the queue is full
func (q *UniqueQueue) isFull() bool {
	return q.length == len(q.queue)
}

// isEmpty checks if the queue is empty
func (q *UniqueQueue) isEmpty() bool {
	return q.length == 0
}

// Push insert a new item to the queue
func (q *UniqueQueue) Push(key interface{}) {
	// check if the queue is full
	var isAvailable bool = !q.isFull()

	// if the queue is not full
	// insert a new item
	if isAvailable {
		// increase the size of queue
		q.length = q.length + 1

		// define the index for the item
		// the item will be inserted to the rear of the queue
		q.rear = 1 + q.rear
		q.queue[q.rear] = key
	} else {
		// remove the oldest data inside the queue
		q.Pop()

		// increase the size of queue
		q.length = q.length + 1

		// define the index for the item
		// the item will be inserted to the rear of the queue
		q.rear = 1 + q.rear
		q.queue[q.rear] = key
	}
}

// Pop returns the first item from the queue
func (q *UniqueQueue) Pop() interface{} {
	// check if the queue is empty
	if q.isEmpty() {
		// if the queue is empty, return empty value
		return nil
	}

	// define the temporary variable called "tmp"
	// this variable contains the first item that will be returned
	var tmp interface{} = q.queue[q.front]

	// iterate through the item inside the queue
	for i := 0; i < q.length-1; i++ {
		// move forward the item
		q.queue[i] = q.queue[i+1]
	}

	// decrease the size of the queue
	q.length = q.length - 1

	// remove the item inside the queue
	q.queue[q.length] = nil
	q.rear--

	// return the first item inside the queue
	return tmp
}

// Contains checks if the item is exists inside the queue
func (q *UniqueQueue) Contains(key interface{}) bool {
	// Iterate through every item in a queue
	for _, val := range q.queue {
		// if the item is found
		// return true
		if val == key {
			return true
		}
	}

	// if the item is not found, return false
	return false
}

// Len returns the size of the queue
func (q *UniqueQueue) Len() int {
	return q.length
}

// Keys returns all the items from the queue
func (q *UniqueQueue) Keys() []interface{} {
	// create a variable to store the items
	var keys []interface{} = []interface{}{}

	// iterate through every item inside the queue
	for i := 0; i < q.length; i++ {
		// insert the value from the queue
		// to the "keys" variable
		keys = append(keys, q.queue[i])
	}

	// return the items from the queue
	return keys
}
