package hash

import (
	lls "dsa_demo/linked-list"
)

type ChainingHashTable struct {
	key   [hashTableSize]lls.LinkList
	value [hashTableSize]lls.LinkList
}

func TestChainingHashTable() {
	ht := initChainingHashTable()
	ht.insert(1234, "Nguyen Van Son")
	ht.insert(3324, "Son Nguyen Van")
}

func initChainingHashTable() ChainingHashTable {
	cnHashTb := ChainingHashTable{}
	for i := 0; i < hashTableSize; i++ {
		cnHashTb.key[i] = lls.LinkList{}

		cnHashTb.value[i] = lls.LinkList{}
	}
	return cnHashTb
}

func (ht *ChainingHashTable) insert(key interface{}, value interface{}) {
	index := HashKey(key)

	ht.key[index].InsertAtTheEnd(value)
}
