package hash

import (
	lls "dsa_demo/linked-list"
	"fmt"
)

const BUCKET int = 10

type ChainingHashTable[T comparable] struct {
	values [BUCKET]lls.GLinkList[T]
}

func TestChainingHashTable() {
	ht := InitChainingHashTable()
	ht.Insert(112, "Son")
	ht.Insert(113, "Tuan")
	ht.Insert(114, "Tung")
	ht.Insert(115, "Manh")
	ht.Insert(116, "Duc")
	ht.Insert(117, "Mai")
	ht.Insert(118, "Ngoc")
	ht.Insert(119, "Bach")
	ht.Insert(120, "Lam")
	ht.Insert(121, "Loc")
	ht.Display()
	ht.GetValue(116)
}

func InitChainingHashTable() ChainingHashTable[int] {
	cnHashTb := ChainingHashTable[int]{}
	for i := 0; i < BUCKET; i++ {
		cnHashTb.values[i] = lls.GLinkList[int]{}
	}
	return cnHashTb
}

func (ht *ChainingHashTable[T]) Insert(key T, value string) {
	index := HashKey(key)
	ht.values[index].InsertAtTheFront(key, value)
}

func (ht *ChainingHashTable[T]) GetValue(key T) interface{} {
	index := HashKey(key)
	node := ht.values[index].SearchNodeByKey(key)
	if node == nil {
		return nil
	}
	fmt.Printf("\nKey: %v - Value: %v", node.Key, node.NodeValue)
	return node
}

func (ht *ChainingHashTable[T]) Display() {
	for i := 0; i < BUCKET; i++ {
		fmt.Printf("\n%d:", i)
		crrNode := ht.values[i].Head
		for crrNode != nil {
			fmt.Printf(" --> %v", crrNode.NodeValue)
			crrNode = crrNode.Next
		}
	}
}
