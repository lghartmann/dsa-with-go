package datastructures

import "fmt"

type HashNode struct {
	key   int
	value string
}
type HashTable struct {
	buckets [100]*LinkedList
	size    int
}

func NewHashTable() *HashTable {
	ht := &HashTable{}
	for i := range ht.buckets {
		ht.buckets[i] = &LinkedList{}
	}
	return ht
}

func (h *HashTable) hashKey(key int) int {
	return key % 100
}

func (h *HashTable) Put(key int, value string) {
	index := h.hashKey(key)
	list := h.buckets[index]

	current := list.head
	for current != nil {
		if node, ok := current.value.(HashNode); ok && node.key == key {
			current.value = HashNode{key: key, value: value}
			return
		}
		current = current.next
	}

	list.Append(HashNode{key: key, value: value})
	h.size++
}

func (h *HashTable) Get(key int) (string, bool) {
	index := h.hashKey(key)
	list := h.buckets[index]

	current := list.head
	for current != nil {
		if node, ok := current.value.(HashNode); ok && node.key == key {
			return node.value, true
		}
		current = current.next
	}
	return "", false
}

func (h *HashTable) Delete(key int) bool {
	index := h.hashKey(key)
	list := h.buckets[index]

	if list.head == nil {
		return false
	}

	current := list.head
	for current != nil {
		if node, ok := current.value.(HashNode); ok && node.key == key {
			list.Delete(current.value)
			h.size--
			return true
		}
		current = current.next
	}
	return false
}

func (h *HashTable) Size() int {
	return h.size
}

func DemonstrateHashTable() {
	ht := NewHashTable()

	ht.Put(1, "One")
	ht.Put(2, "Two")
	ht.Put(101, "One Hundred One")

	if val, ok := ht.Get(1); ok {
		fmt.Printf("Key 1: %s\n", val)
	}
	if val, ok := ht.Get(101); ok {
		fmt.Printf("Key 101: %s\n", val)
	}

	ht.Delete(2)

	fmt.Printf("Hash table size: %d\n", ht.Size())
}
