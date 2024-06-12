package heap

import (
	"fmt"
	"math"
)

/*
Links:
https://www.geeksforgeeks.org/heap-data-structure/
https://www.geeksforgeeks.org/introduction-to-heap-data-structure-and-algorithm-tutorials/
*/

type MaxHeap struct {
	arr  []int
	max  int
	size int
}

func RunTest() {
	maxHeap := initMaxHeap(10)
	maxHeap.insertKey(1)
	maxHeap.insertKey(5)
	maxHeap.insertKey(4)
	maxHeap.insertKey(2)
	maxHeap.insertKey(8)
	maxHeap.insertKey(7)
	maxHeap.insertKey(3)
	maxHeap.insertKey(3)
	//maxHeap.removeMax()
	fmt.Println(maxHeap.arr)
	maxHeap.increaseKey(4, 30)
	fmt.Println("After increaseKey:")
	fmt.Println(maxHeap.arr)

	maxHeap.removeMax()
	fmt.Println("After Remove Max:")
	fmt.Println(maxHeap.arr)

	maxHeap.deleteKey(3)
	fmt.Println("After Remove Key:")
	fmt.Println(maxHeap.arr)
}

func initMaxHeap(maxSize int) MaxHeap {
	maxHeap := MaxHeap{
		max:  maxSize,
		size: 0,
		arr:  make([]int, maxSize),
	}
	return maxHeap
}

func (heap *MaxHeap) maxHeapify(i int) {
	l := lChild(i)
	r := rChild(i)
	largest := i
	if l < heap.size && heap.arr[l] > heap.arr[i] {
		largest = l
	}
	if r < heap.size && heap.arr[r] > heap.arr[largest] {
		largest = r
	}
	if largest != i {
		heap.arr[i], heap.arr[largest] = heap.arr[largest], heap.arr[i]
		heap.maxHeapify(largest)
	}
}

func (heap *MaxHeap) insertKey(x int) {
	if heap.size == heap.max {
		fmt.Println("Heap overflow!!!")
		return
	}

	heap.size++
	i := heap.size - 1
	heap.arr[i] = x

	// check max -> restored if violation
	for i != 0 && heap.arr[parent(i)] < heap.arr[i] {
		heap.arr[i], heap.arr[parent(i)] = heap.arr[parent(i)], heap.arr[i]
		i = parent(i)
	}
}

func (heap *MaxHeap) removeMax() int {
	if heap.size <= 0 {
		return 0
	}
	if heap.size == 1 {
		heap.size--
		return heap.arr[0]
	}

	root := heap.arr[0]
	heap.arr[0], heap.arr[heap.size-1] = heap.arr[heap.size-1], 0
	heap.size--
	heap.maxHeapify(0)
	return root
}

func (heap *MaxHeap) increaseKey(i int, newValue int) {
	heap.arr[i] = newValue
	for i != 0 && heap.arr[parent(i)] < heap.arr[i] {
		heap.arr[i], heap.arr[parent(i)] = heap.arr[parent(i)], heap.arr[i]
		i = parent(i)
	}
}

func (heap *MaxHeap) deleteKey(i int) {
	heap.increaseKey(i, math.MaxUint32)
	heap.removeMax()
}

func lChild(i int) int {
	return (2*i + 1)
}

func rChild(i int) int {
	return (2*i + 2)
}

func parent(i int) int {
	return (i - 1) / 2
}
