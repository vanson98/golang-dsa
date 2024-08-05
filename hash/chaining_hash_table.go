package hash

import (
	lls "dsa_demo/linked-list"
)

type ChainingHashTable struct {
	values [hashTableSize]lls.LinkList
}

func TestChainingHashTable() {
	ht := InitChainingHashTable()
	ht.Insert(1234, "Nguyen Van Son")
	ht.Insert(3324, "Son Nguyen Van")
}

func InitChainingHashTable() ChainingHashTable {
	cnHashTb := ChainingHashTable{}
	for i := 0; i < hashTableSize; i++ {
		cnHashTb.values[i] = lls.LinkList{}
	}
	return cnHashTb
}

func (ht *ChainingHashTable) Insert(key int, value string) {
	index := HashKey(key)
	ht.values[index].InsertAtTheEnd(value)
}

func (ht *ChainingHashTable) GetValue(key int) {
	//index := HashKey(key)

	//ht.values[index].SearchByValue()
}
