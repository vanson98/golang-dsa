package hash

func TestStringHashing() {
	ht := HashTableLinear{
		size: 10,
	}
	ht.insert("K001", "Nguyen Van Son 0")
	ht.insert("K002", "Nguyen Van Son 1")
	ht.insert("K003", "Nguyen Van Son 2")
	ht.insert("K004", "Nguyen Van Son 3")
}

func (ht *HashTableLinear) hashString(str string) int {
	var sum int
	for _, v := range str {
		sum += int(v)
	}
	return sum / ht.size
}
