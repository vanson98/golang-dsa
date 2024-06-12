package arr

import (
	"fmt"
)

/*
Link: https://www.geeksforgeeks.org/program-to-reverse-an-array/
*/

func ReverseArray(arr []int) {
	arrLength := len(arr)
	reverseArr := make([]int, arrLength)
	for i, _ := range arr {
		reverseArr[i] = arr[arrLength-i-1]
	}
	for _, v := range reverseArr {
		fmt.Println(v)
	}
}
