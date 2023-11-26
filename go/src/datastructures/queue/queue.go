package queue

/*
Queue is an implementation of a batched queue
*/
type Queue[T any] struct {
	front *node[T]
	back  *node[T]
}

type node[T any] struct {
	val    T
	next   *node[T]
	length int
}

/*
New creates a new Queue
*/
func New[T any]() *Queue[T] {
	return &Queue[T]{nil, nil}
}

/*
Enqueue adds a val to the Queue
Complexity: O(1)
*/
func (q *Queue[T]) Enqueue(val T) {
	if q.front == nil {
		q.front = &node[T]{val, nil, 1}
	} else {
		if q.back == nil {
			q.back = &node[T]{val, nil, 1}
		} else {
			newNode := &node[T]{val, q.back, q.back.length + 1}
			q.back = newNode
		}
	}
}

/*
Dequeue removes the first element in the Queue
Complexity: Amortized O(1)
*/
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.front == nil {
		var empty T
		return empty, false
	} else if q.front.next == nil {
		front := q.front
		q.front = q.back.reverse()
		q.back = nil
		return front.val, true
	} else {
		front := q.front
		q.front = q.front.next
		return front.val, true
	}
}

/*
Len returns the length of the Queue
Complexity: O(1)
*/
func (q Queue[T]) Len() int {
	var length int
	if q.front != nil {
		length += q.front.length
	}
	if q.back != nil {
		length += q.back.length
	}
	return length
}

func (n *node[T]) reverse() *node[T] {
	var prev *node[T]
	curr := n
	length := 0
	for curr != nil {
		next := curr.next
		curr.next = prev
		length++
		curr.length = length
		prev, curr = curr, next
	}
	return prev
}

/*
Peek returns the first element in the Queue without removing it
*/
func (q Queue[T]) Peek() (T, bool) {
	if q.front == nil {
		var empty T
		return empty, false
	} else {
		return q.front.val, true
	}
}
