package datastructures

import "fmt"

type DoublyLinkedList struct {
	head   *DoublyLinkedListNode
	length int
}

type DoublyLinkedListNode struct {
	prev  *DoublyLinkedListNode
	value int
	next  *DoublyLinkedListNode
}

func (l *DoublyLinkedList) Length() int {
	return l.length
}

func (l *DoublyLinkedList) Prepend(val int) {
	newNode := &DoublyLinkedListNode{
		prev:  nil,
		value: val,
		next:  l.head,
	}
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	l.length++
	fmt.Println("New head:", l.head.value)
}

func (l *DoublyLinkedList) Append(val int) {
	newNode := &DoublyLinkedListNode{
		prev:  nil,
		value: val,
		next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		l.length++
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
	l.length++
	fmt.Println("DoublyHead remains:", l.head.value)
}

func (l *DoublyLinkedList) Delete(val int) {
	if l.length <= 1 {
		fmt.Println("Your linked list has less than 2 nodes")
		return
	}
	current := l.head
	for current != nil {
		if current.value == val {
			if current.prev == nil {
				l.head = current.next
				if l.head != nil {
					l.head.prev = nil
				}
			} else if current.next == nil {
				current.prev.next = nil
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			l.length--
			return
		}
		current = current.next
	}
}

func (l *DoublyLinkedList) List() {
	if l.length == 0 {
		fmt.Println("No nodes in your list")
		return
	}
	current := l.head
	for current != nil {
		fmt.Println("Value:", current.value, "Prev:", current.prev, "Next:", current.next)
		current = current.next
	}
}

func (l *DoublyLinkedList) Clear() {
	l.head = nil
	l.length = 0
}

func DemonstrateDoublyLinkedList() {
	list := &DoublyLinkedList{}
	list.Prepend(5)
	list.Append(6)
	list.Prepend(7)
	list.Append(8)
	list.Prepend(9)
	list.Append(10)
	list.Prepend(1)

	list.List()
	list.Delete(1)
	list.Delete(6)
	list.List()
	fmt.Println("Length:", list.Length())
	list.Clear()
	fmt.Println("Cleared List:", list.head)
}
