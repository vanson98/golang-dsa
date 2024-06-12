package arr

import (
	"fmt"
)

/*
Links: https://www.geeksforgeeks.org/rearrange-array-arri/
*/

// Complexity: O(n2)
func Rearrange(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j, _ := range arr {
			if arr[j] == i {
				arr[j] = arr[i]
				arr[i] = i
				break
			}
		}
	}
	for _, v := range arr {
		fmt.Print(fmt.Sprintf("%v, ", v))
	}
}

// Complexity: O(n)
func Rearrange2(arr []int) {
	for i, v := range arr {
		if v == -1 || v == i {
			continue
		} else {
			x := arr[i]
			arr[i] = arr[v]
			arr[x] = x
		}
	}

	for _, v := range arr {
		fmt.Print(fmt.Sprintf("%v, ", v))
	}
}

func Rearrange3(a []int) {
	for i, _ := range a {
		if a[i] != -1 && a[i] != i {
			var x int = a[i]
			for a[x] != -1 && a[x] != x {
				var y int = a[x]
				a[x] = x
				x = y
			}
			a[x] = x
			if a[i] != i {
				a[i] = -1
			}
		}
	}
	for _, v := range a {
		fmt.Print(fmt.Sprintf("%v, ", v))
	}
}

func Rearrange4(a []int) {
	//n := len(a)
	s := make(map[int]bool)
	for _, v := range a {
		if v > -1 {
			s[v] = true
		}
	}

	for i, _ := range a {
		if _, ok := s[i]; ok {
			a[i] = i
		} else {
			a[i] = -1
		}
	}
	fmt.Print(a)
}
