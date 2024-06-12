package hash

import (
	"fmt"
)

/*
Link: https://www.tutorialspoint.com/golang-program-to-implement-a-hash-table-with-linear-probing
*/
const hashTableSize = 20

type HashTableLinear struct {
	keys   [hashTableSize]interface{}
	values [hashTableSize]interface{}
	size   int
}

func (ht *HashTableLinear) hashByDevision(key int) int {
	return key % ht.size
}

func (ht *HashTableLinear) hashKey(key interface{}) int {
	var index int
	switch value := key.(type) {
	case int:
		index = ht.hashByDevision(value)
	case string:
		index = ht.hashString(value)
	}
	return index
}

func (ht *HashTableLinear) insert(key interface{}, value interface{}) {
	index := ht.hashKey(key)

	// check duplicate key
	if ht.keys[index] == key {
		ht.values[index] = value
		return
	}

	for ht.keys[index] != nil {
		index = (index + 1) % ht.size
	}

	ht.keys[index] = key
	ht.values[index] = value
}

func (ht *HashTableLinear) get(key interface{}) interface{} {
	index := ht.hashKey(key)

	for ht.keys[index] != key {
		index = (index + 1) % hashTableSize
	}

	return ht.values[index]
}

func TestHashTableLinear() {
	ht := HashTableLinear{
		size: hashTableSize,
	}

	ht.insert(234, "Son")
	ht.insert(10, "Tien")
	ht.insert(10, "Tien Bip")
	ht.insert(33, "Tung")
	ht.insert(43, "Chung")

	ht.insert("K001", "Nguyen Van Son 0")
	ht.insert("K002", "Nguyen Van Son 1")
	ht.insert("K003", "Nguyen Van Son 2")
	ht.insert("K004", "Nguyen Van Son 3")

	fmt.Println(ht.get("K002"))
	// fmt.Println(ht.get(234))
	// fmt.Println(ht.get(43))
	// fmt.Println(ht.get(33))

	fmt.Println(ht.keys[:]...)
	fmt.Println(ht.values[:]...)
}
