package mtx

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test() {
	//initMatrix(3,4)
	//searchInMatrix(9)
	//searchInMatrix2(28)
	//searchInMatrix3(128)
	searchInMatrix4(50)
}

func printMatrix(matrix [][]int) {
	for _, r := range matrix {
		for _, ri := range r {
			fmt.Printf("%d\t", ri)
		}
		fmt.Println()
	}
}

func initMatrix(row uint, col uint) {
	matrix := make([][]int, row)
	reader := bufio.NewReader(os.Stdin)
	for i, _ := range matrix {
		matrix[i] = make([]int, col)
		for j, _ := range matrix[i] {
			fmt.Printf("Value at matrix[%d][%d]: ", i, j)
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			value, err := strconv.Atoi(text)
			if err != nil {
				value = 0
			}
			matrix[i][j] = value
		}
	}
	printMatrix(matrix)
}

// Brute force
func searchInMatrix(x int) bool {
	matrix := [][]int{
		{1, 5, 9},
		{14, 20, 21},
		{30, 34, 43},
	}
	for _, row := range matrix {
		for _, ri := range row {
			if ri == x {
				fmt.Println("Found")
				return true
			}
		}
	}
	fmt.Println("Not Found")
	return false
}

// Typecasted array
func searchInMatrix2(x int) bool {
	matrix := [][]int{
		{0, 6, 8, 9, 11},
		{20, 22, 28, 29, 31},
		{36, 38, 50, 61, 63},
		{64, 66, 100, 122, 128},
	}
	// cast to 1D array
	l := 0
	for _, row := range matrix {
		l += len(row)
	}

	arr := make([]int, 0, l)
	for _, v := range matrix {
		arr = append(arr, v...)
	}
	for _, v := range arr {
		if v == x {
			fmt.Print("Found")
			return true
		}
	}
	fmt.Print("Not Found")
	return true
}

// Binary search
func searchInMatrix3(x int) {
	mat := [][]int{
		{0, 6, 8, 9, 11},
		{20, 22, 28, 29, 31},
		{36, 38, 50, 61, 63},
		{64, 66, 100, 122, 128},
	}
	n := 4
	m := 5
	/*
	   search on middle column
	   till only two elements are left
	   or middle element of some row in the search is the required element x
	   --> skip the rows that are not required
	*/

	i_low := 0
	i_hight := n - 1
	j_mid := m / 2

	for i_low+1 < i_hight {
		i_mid := (i_hight + i_low) / 2

		if mat[i_mid][j_mid] == x {
			fmt.Printf("found at (%v,%v)", i_mid, j_mid)
			return
		} else if mat[i_mid][j_mid] > x {
			i_hight = i_mid
		} else {
			i_low = i_mid
		}
	}
	/*
	   two left element must be adjacent
	   whether x equals to the middle element
	   otherwise according to the value of x check
	       present in the
	           1st half of 1st row,
	           2nd half of 1st row,
	           1st half of 2nd row,
	           2nd half of 2nd row
	*/

	if mat[i_low][j_mid] == x {
		fmt.Printf("found at (%v,%v)", i_low, j_mid)
		return
	} else if mat[i_hight][j_mid] == x {
		fmt.Printf("found at (%v,%v)", i_hight, j_mid)
		return
	} else if x <= mat[i_low][j_mid-1] {
		binarySearch(x, i_low, 0, j_mid-1, mat)
	} else if mat[i_low][j_mid+1] <= x && x <= mat[i_low][m-1] {
		binarySearch(x, i_low, j_mid+1, m-1, mat)
	} else if mat[i_hight][0] <= x && x <= mat[i_hight][j_mid] {
		binarySearch(x, i_hight, 0, j_mid-1, mat)
	} else if mat[i_hight][j_mid+1] <= x && x <= mat[i_hight][m-1] {
		binarySearch(x, i_hight, j_mid+1, m-1, mat)
	}
}

// Binary search 2
func searchInMatrix4(x int) {
	mat := [][]int{
		{0, 6, 8, 9, 11},
		{20, 22, 28, 29, 31},
		{36, 38, 50, 61, 63},
		{64, 66, 100, 122, 128},
	}
	n := 4
	m := 5
	low := 0
	hight := n - 1
	// find row
	for low <= hight {
		mid := (low + hight) / 2
		if x < mat[mid][0] {
			hight = mid - 1
		} else if x > mat[mid][m-1] {
			low = mid + 1
		} else if x > mat[mid][0] && x < mat[mid][m-1] {
			binarySearch(x, mid, 0, m-1, mat)
			return
		} else if x == mat[mid][0] {
			fmt.Printf("Found at mat[%v][%v]", mid, 0)
			return
		} else {
			fmt.Printf("Found at mat[%v][%v]", mid, m-1)
			return
		}
	}
	fmt.Println("Not Found")
}

func binarySearch(x int, i int, j_low int, j_hight int, mat [][]int) {
	for j_low <= j_hight {
		mid := (j_low + j_hight) / 2
		if mat[i][mid] == x {
			fmt.Printf("found at (%v,%v)", i, mid)
			return
		} else if x < mat[i][mid] {
			j_hight = mid - 1
		} else {
			j_low = mid + 1
		}
	}
	fmt.Println("Not Found")
}
