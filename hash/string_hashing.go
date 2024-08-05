package hash

func TestStringHashing() {
	ht := HashTableLinear{
		size: 10,
	}
	ht.Insert("K001", "Nguyen Van Son 0")
	ht.Insert("K002", "Nguyen Van Son 1")
	ht.Insert("K003", "Nguyen Van Son 2")
	ht.Insert("K004", "Nguyen Van Son 3")
}

func (ht *HashTableLinear) hashString(str string) int {
	var sum int
	for _, v := range str {
		sum += int(v)
	}
	return sum / ht.size
}

func stringFoldingHashing(s string) int {
	var sum int = 0
	var mul int = 1
	for i := 0; i < len(s); i++ {
		if i%4 == 0 {
			mul = 1
		} else {
			mul *= 256
		}
		sum = sum + (int(s[i]) * mul)
	}
	return sum % hashTableSize
}
