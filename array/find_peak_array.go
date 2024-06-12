package arr

import (
	"fmt"
)

/*
Link: https://www.geeksforgeeks.org/find-a-peak-in-a-given-array/
*/
func FindPeekElement(arr []int) {
	peekResult := make([]int, 0)
	n := len(arr)
	for i, v := range arr {
		if (i > 0 && i < n-1) && (arr[i] > arr[i-1] && arr[i] > arr[i+1]) {
			peekResult = append(peekResult, v)
		}
	}
	if len(peekResult) == 0 && arr[0] >= arr[1] {
		peekResult = append(peekResult, arr[0])
	} else if len(peekResult) == 0 && arr[n-1] >= arr[n-2] {
		peekResult = append(peekResult, arr[n-1])
	}
	fmt.Printf("\n All peak: %v", peekResult)
}

func FindPeekBinarySearch() {
	//arr := []int{231,80,69,54,47,39,26,15,1}
	arr := []int{1, 15, 25, 34, 45, 5555, 58, 50, 20, 10}
	fmt.Println(FindPeekBNS1(arr))
	fmt.Println(FindPeekBNS2(arr, 0, len(arr)-1))
}

func FindPeekBNS1(arr []int) int {
	var peak int
	low := 0
	high := len(arr) - 1
	mid := (low + high) / 2

	if mid < high && arr[mid] < arr[mid+1] {
		low = mid + 1
		peak = FindPeekBNS1(arr[low:])
	} else if mid > 0 && arr[mid-1] > arr[mid] {
		peak = FindPeekBNS1(arr[:mid])
	} else {
		peak = arr[mid]
	}
	return peak
}

func FindPeekBNS2(arr []int, low int, high int) int {
	mid := (low + high) / 2
	if mid < high && arr[mid] < arr[mid+1] {
		low = mid + 1
		return FindPeekBNS2(arr, low, high)
	} else if mid > 0 && arr[mid-1] > arr[mid] {
		high = mid - 1
		return FindPeekBNS2(arr, low, high)
	} else {
		return arr[mid]
	}
}
