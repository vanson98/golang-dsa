package hash

import (
	"fmt"
)

func TestHashMapTable() {
	ht := make(map[int]interface{})

	ht[311] = 66666
	ht[23] = "Son"
	ht[1] = "Tien"

	if value, ok := ht[311]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exits in hash table")
	}

	//delete(ht,1)

	if value, ok := ht[1]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exits in hash table")
	}
}
