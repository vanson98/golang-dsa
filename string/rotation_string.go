package str

import (
	"fmt"
)

// Method 0: Loop and apprend

func LeftRotateString(str string, n int) {
	strLen := len(str)
	tempStr := make([]byte, strLen)
	for i, _ := range str {
		if i+n >= strLen {
			tempStr = append(tempStr, str[:n]...)
			break
		}
		tempStr[i] = str[i+n]
	}
	fmt.Println(string(tempStr))
}

func RightRotateString(str string, n int) {
	strLen := len(str)

	//byteString := []byte(str)
	byteString := make([]byte, strLen)
	copy(byteString, str)

	tempStr := make([]byte, strLen)

	for i := strLen - 1; i >= 0; i-- {
		if i-n < 0 {
			tempStr = append(byteString[strLen-n:], tempStr...)
			break
		}
		tempStr[i] = str[i-n]
	}

	fmt.Println(string(tempStr))
}

// Method 1: Join two string

func LeftRotateString2(str string, n int) {
	subString1 := str[n:]
	subString2 := str[0:n]
	fmt.Println(subString1 + subString2)
}

func RightRotateString2(str string, n int) {
	LeftRotateString(str, len(str)-n)
}

// Method 2: Use extended string

func LeftRotateString3(str string, n int) {
	tempStr := str + str
	rotateString := tempStr[n : len(str)+n]
	fmt.Println(rotateString)
}

func RightRotateString3(str string, n int) {
	LeftRotateString3(str, len(str)-n)
}

// Method 3: Use queue
func LeftRotateString4(str string, n int) {
	queue := []rune(str)
	//copy(queue,str)

	for i := 0; i < n; i++ {
		var character rune
		character, queue = dequeue(queue)
		queue = enqueue(queue, character)
	}

	fmt.Println(string(queue))
}

func enqueue(queue []rune, element rune) []rune {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []rune) (rune, []rune) {
	element := queue[0]
	if len(queue) == 1 {
		var tmp = []rune{}
		return element, tmp
	}

	return element, queue[1:]
}
