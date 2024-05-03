package bif

import (
	"fmt"
)

func Foo(is ...int) {
	for i := 0; i < len(is); i++ {
		fmt.Println(is[i])
	}
}
