package hash

import (
	"fmt"
	"math"
	"strconv"
)

/*
Link:
https://www.tutorialspoint.com/golang-program-to-implement-a-hash-table-with-linear-probing
https://medium.com/enjoy-algorithm/introduction-to-hashing-in-programming-617960aeccf2
*/
const hashTableSize = 50

type HashTableLinear struct {
	keys   [hashTableSize]interface{}
	values [hashTableSize]interface{}
	size   int
}

func devisionHashing(key int) int {
	return key % hashTableSize
}

func midSquareHashing(key int) int {
	sq := key * key
	sq_string := strconv.Itoa(sq)
	hashtb_range := len(strconv.Itoa(hashTableSize))
	sq_len := len(sq_string)
	mid_digits_index := (sq_len / 2) - 1
	mid_digits := sq_string[mid_digits_index : mid_digits_index+hashtb_range]
	result, _ := strconv.Atoi(mid_digits)
	if result > hashTableSize {
		return devisionHashing(result)
	}
	return result
}

func mutiplicationHashing(key int) int {
	var multiplicationConst float64 = ((math.Sqrt(5) - 1) / 2)
	fractionalPath := float64(key)*multiplicationConst - math.Floor(float64(key)*multiplicationConst)
	hashValue := math.Floor(fractionalPath * float64(hashTableSize))
	return int(hashValue)
}

func HashKey(key interface{}) int {
	var index int
	switch value := key.(type) {
	case int:
		//index = ht.hashByDevision(value)
		//index = ht.hashByMidSquare(value)
		index = mutiplicationHashing(value)
	case string:
		//index = ht.hashString(value)
		index = stringFoldingHashing(value)
	}
	return index
}

func (ht *HashTableLinear) insert(key interface{}, value interface{}) {
	index := HashKey(key)

	// check duplicate key
	if ht.keys[index] == key {
		ht.values[index] = value
		return
	}

	for ht.keys[index] != nil {
		fmt.Println("Collision")
		index = (index + 1) % ht.size
	}

	ht.keys[index] = key
	ht.values[index] = value
}

func (ht *HashTableLinear) get(key interface{}) interface{} {
	index := HashKey(key)

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
	ht.insert(123, "Son")
	ht.insert(3245345, "Tien")
	ht.insert(234, "Tien Bip")
	ht.insert(768, "Tung")
	ht.insert(433333, "Chung")

	ht.insert(431, "Chung")
	ht.insert(1231, "Son")
	ht.insert(32453453, "Tien")
	ht.insert(2345, "Tien Bip")
	ht.insert(76899, "Tung")
	ht.insert(4333, "Chung")

	// ht.insert("K001", "Nguyen Van Son 0")
	// ht.insert("K002", "Nguyen Van Son 1")
	// ht.insert("K003", "Nguyen Van Son 2")
	// ht.insert("K004", "Nguyen Van Son 3")

	//fmt.Println(ht.get("K002"))
	// fmt.Println(ht.get(234))
	fmt.Println(ht.get(43))
	// fmt.Println(ht.get(33))

	fmt.Println(ht.keys[:]...)
	fmt.Println(ht.values[:]...)
}
