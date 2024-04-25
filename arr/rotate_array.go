package arr

import (
	"fmt"
)

/*

Link: https://www.geeksforgeeks.org/c-program-cyclically-rotate-array-one/

*/

func RotateArray(arr []int) {
	arrLength := len(arr)
	lastElement := arr[arrLength-1]
	for i := arrLength - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = lastElement
	for _, v := range arr {
		fmt.Print(v)
	}
}

func RotateArrayWithTwoPointers(arr []int) {
	n := len(arr)
	i, j := 0, n-1
	for i != j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
	}
	for _, v := range arr {
		fmt.Print(v)
	}
}
