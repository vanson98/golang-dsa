package heap

import (
	"fmt"
	"math/rand"
)

func TestHeapSort() {
	arrSize := 20
	arr := make([]int, 0)

	for i := 0; i < arrSize; arrSize-- {
		arr = append(arr, rand.Intn(arrSize))
	}

	arrLength := len(arr)
	maxHeap := MaxHeap{
		max:  arrLength,
		size: 0,
		arr:  make([]int, arrLength),
	}
	for _, v := range arr {
		maxHeap.insertKey(v)
	}
	fmt.Println(maxHeap.arr)

	sortedArr := make([]int, arrLength)
	for maxHeap.size > 0 {
		max := maxHeap.removeMax()
		sortedArr[maxHeap.size] = max
	}

	fmt.Println(sortedArr)
}
