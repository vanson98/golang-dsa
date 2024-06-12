package str

import (
	"fmt"
	"strings"
)

func PrintAllSubStrings(str string) {
	n := len(str)
	for i, _ := range str {
		temp := make([]byte, n-i+1)
		tempindex := 0
		for j := i; j < n; j++ {
			temp[tempindex] = str[j]
			tempindex++
			temp[tempindex] = 0
			fmt.Println(string(temp[:]))
		}
	}
}

func CountNumberOfSubstrings(s1 string, s2 string) {
	n := len(s1)
	count := 0
	var subStr string
	for i, _ := range s1 {
		temp := make([]byte, 0, 0)
		for j := i; j < n; j++ {
			temp = append(temp, s1[j])
			subStr = string(temp[:])
			if isContain := strings.Contains(s2, subStr); isContain {
				count++
			}
		}
	}
	fmt.Printf("\nTotal substring is contained in str2: %d", count)
}
