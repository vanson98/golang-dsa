package queue

import (
	"fmt"
)

func TestQueue() {
	myqueue := Queue[int]{cap: 5, rear: -1, front: -1}
	myqueue.initQueue()

	myqueue.enqueue(2)
	myqueue.enqueue(3)
	myqueue.enqueue(6)
	myqueue.enqueue(6)
	myqueue.enqueue(6)

	myqueue.isFull()

	myqueue.dequeue()
	myqueue.dequeue()
	myqueue.dequeue()
	myqueue.dequeue()
	myqueue.dequeue()
	myqueue.isEmpty()

	myqueue.enqueue(62)
	myqueue.enqueue(16)
	myqueue.enqueue(61)
	myqueue.enqueue(33)
	fmt.Print(myqueue.data)
}

type Queue[T any] struct {
	data  []T
	cap   int
	rear  int
	front int
}

func (q *Queue[T]) initQueue() {
	q.data = make([]T, q.cap)
}

func (q *Queue[T]) enqueue(element T) {
	if q.rear == q.cap-1 {
		panic("Queue is full")
	} else {
		if q.front == -1 {
			q.front = 0
		}
		q.rear++
		q.data[q.rear] = element
	}
}

func (q *Queue[T]) dequeue() T {
	if q.isEmpty() {
		return getZero[T]()
	} else {
		elm := q.data[q.front]
		for i := 0; i < q.rear; i++ {
			q.data[i] = q.data[i+1]
		}
		q.data[q.rear] = getZero[T]()
		q.rear--
		if q.rear == -1 {
			q.front = -1
		}
		return elm
	}
}

func (q *Queue[T]) isEmpty() bool {
	if q.rear == -1 && q.front == -1 {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (q *Queue[T]) isFull() bool {
	if q.rear == q.cap-1 && q.front == 0 {
		fmt.Println("Queue is full")
		return true
	}
	return false
}

func (q *Queue[T]) getFront() T {
	if q.isEmpty() {
		return getZero[T]()
	}
	return q.data[q.front]
}

func (q *Queue[T]) getRear() T {
	if q.isEmpty() {
		return getZero[T]()
	}
	return q.data[q.rear]
}

func getZero[T any]() T {
	var result T
	return result
}
