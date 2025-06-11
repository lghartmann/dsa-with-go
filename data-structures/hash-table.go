package datastructures

type HashTable map[int]string

type HashLinkedList struct {
	val  string
	next *HashLinkedList
}

func (l *HashLinkedList) append(val string) {

}

var Hash HashTable = make(map[int]string)

func (h *HashTable) addItem(key int, val string) {
	hashedKey := h.hashKey(key)

}

func (h *HashTable) hashKey(val int) int {
	return val % 100
}
