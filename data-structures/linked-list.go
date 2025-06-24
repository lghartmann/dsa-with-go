package datastructures

import "fmt"

type LinkedList struct {
	head   *LinkedListNode
	length int
}

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Prepend(val interface{}) {
	newNode := &LinkedListNode{
		value: val,
		next:  l.head,
	}
	l.head = newNode
	l.length++
}

func (l *LinkedList) Append(val interface{}) {
	newNode := &LinkedListNode{
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
	l.length++
}

func (l *LinkedList) Delete(val interface{}) {
	if l.head == nil {
		return
	}
	if l.head.value == val {
		l.head = l.head.next
		l.length--
		return
	}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current.value == val {
			prev.next = current.next
			l.length--
			return
		}
		prev = current
		current = current.next
	}
}

func (l *LinkedList) List() {
	if l.length == 0 {
		fmt.Println("No nodes in your list")
		return
	}
	current := l.head
	for current != nil {
		fmt.Println("Value:", current.value, "Next:", current.next)
		current = current.next
	}
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.length = 0
}

func DemonstrateLinkedList() {
	list := &LinkedList{}
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
	fmt.Println(list.Length())
	list.Clear()
	fmt.Println("Cleared List:", list.head)
}
