package queue

import (
	"fmt"
)

func TestQueue() {
	myqueue := InitQueue[int](100)

	myqueue.Enqueue(2)
	myqueue.Enqueue(3)
	myqueue.Enqueue(6)
	myqueue.Enqueue(6)
	myqueue.Enqueue(6)

	myqueue.IsFull()

	myqueue.Dequeue()
	myqueue.Dequeue()
	myqueue.Dequeue()
	myqueue.Dequeue()
	myqueue.Dequeue()
	myqueue.IsEmpty()

	myqueue.Enqueue(62)
	myqueue.Enqueue(16)
	myqueue.Enqueue(61)
	myqueue.Enqueue(33)
	fmt.Print(myqueue.data)
}

type Queue[T any] struct {
	cap   int
	data  []T
	rear  int
	front int
}

func InitQueue[T any](cap int) Queue[T] {
	return Queue[T]{
		cap:   cap,
		data:  make([]T, cap),
		rear:  -1,
		front: -1,
	}
}

func (q *Queue[T]) Enqueue(element T) {
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

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
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

func (q *Queue[T]) IsEmpty() bool {
	if q.rear == -1 && q.front == -1 {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (q *Queue[T]) IsFull() bool {
	if q.rear == q.cap-1 && q.front == 0 {
		fmt.Println("Queue is full")
		return true
	}
	return false
}

func (q *Queue[T]) GetFront() T {
	if q.IsEmpty() {
		return getZero[T]()
	}
	return q.data[q.front]
}

func (q *Queue[T]) GetRear() T {
	if q.IsEmpty() {
		return getZero[T]()
	}
	return q.data[q.rear]
}

func getZero[T any]() T {
	var result T
	return result
}
