package queue

import (
	"fmt"
	"slices"
)

/*
Implement Deque: https://www.programiz.com/dsa/deque
*/

type Deque[T any] struct {
	data []T
	size int
}

func TestDeque() {
	deque := initDeque[int](6)
	deque.pushFront(10)
	deque.pushFront(11)
	deque.pushFront(12)
	deque.pushFront(13)
	deque.pushFront(14)
	deque.pushFront(15)
	deque.pushFront(15)

	deque.popRear()
	deque.popRear()
	deque.popRear()

	fmt.Println(deque.data)

	deque2 := initDeque[int](6)
	deque2.pushRear(10)
	deque2.pushRear(11)
	deque2.pushRear(12)
	deque2.pushRear(13)
	deque2.pushRear(14)

	deque2.popFront()
	deque2.popFront()
	deque2.popFront()
	deque2.popFront()
	deque2.popFront()
	deque2.popFront()
	fmt.Println(deque2.data)

}

func initDeque[T any](size int) Deque[T] {
	deq := Deque[T]{size: size}
	deq.data = make([]T, 0)
	return deq
}

func (q *Deque[T]) isFull() bool {
	if len(q.data) == q.size {
		fmt.Println("Queue is overflow")
		return true
	}
	return false
}

func (q *Deque[T]) isEmpty() bool {
	if len(q.data) == 0 {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (q *Deque[T]) pushFront(element T) {
	if !q.isFull() {
		q.data = slices.Insert(q.data, 0, element)
	}
}

func (q *Deque[T]) pushRear(elm T) {
	if !q.isFull() {
		q.data = append(q.data, elm)
	}
}

func (q *Deque[T]) popFront() (T, bool) {
	if q.isEmpty() {
		return getZero[T](), false
	}
	front := q.data[0]
	q.data = q.data[1:]
	return front, true
}

func (q *Deque[T]) popRear() (T, bool) {
	if q.isEmpty() {
		return getZero[T](), false
	}
	rearElm := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return rearElm, true
}
