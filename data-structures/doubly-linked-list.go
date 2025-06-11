package datastructures

import "fmt"

var doublyLength = 0
var DoublyHead *DoublyLinkedListNode

type DoublyLinkedListNode struct {
	prev  *DoublyLinkedListNode
	value int
	next  *DoublyLinkedListNode
}

func (n *DoublyLinkedListNode) Length() int {
	return doublyLength
}

func (n *DoublyLinkedListNode) Prepend(val int) {
	if n == nil {
		DoublyHead = &DoublyLinkedListNode{
			prev:  nil,
			value: val,
			next:  nil,
		}
		doublyLength++
		return
	}

	newNode := &DoublyLinkedListNode{
		prev:  nil,
		value: val,
		next:  n,
	}

	n.prev = newNode

	DoublyHead = newNode

	fmt.Println("New head: ", DoublyHead.value)
	doublyLength++
}

func (n *DoublyLinkedListNode) Append(val int) {
	if n == nil {
		DoublyHead = &DoublyLinkedListNode{
			prev:  nil,
			value: val,
			next:  nil,
		}
		doublyLength++
		return
	}

	current := n

	// we have to guarantee the iteration through the end of the doubly linked list
	for current.next != nil {
		current = current.next
	}

	current.next = &DoublyLinkedListNode{
		prev:  current,
		value: val,
		next:  nil,
	}

	doublyLength++
	fmt.Println("DoublyHead remains: ", n.value)
}

func (n *DoublyLinkedListNode) Delete(val int) {
	if n.Length() <= 1 {
		// no node will be deleted if we have a single node in our list
		fmt.Println("Your linked list has less than 2 nodes")
		return
	}

	current := n

	for current != nil {
		if current.value == val {
			if current.prev == nil {
				DoublyHead = current.next
				if DoublyHead != nil {
					DoublyHead.prev = nil
				}
			} else if current.next == nil {
				current.prev.next = nil
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			doublyLength--
		}
		current = current.next
	}
}

func (n *DoublyLinkedListNode) List() {
	if n.Length() == 0 {
		fmt.Println("No nodes in your list")
		return
	}

	current := n

	for current != nil {
		fmt.Println("Value: ", current.value, "Prev: ", current.prev, "Next: ", current.next)
		current = current.next
	}
}

func (n *DoublyLinkedListNode) Clear() {
	fmt.Println("LOL NO NEED TO FREE MEMORY")
	DoublyHead = nil
}

func DemonstrateDoublyLinkedList() {
	DoublyHead.Prepend(5)
	DoublyHead.Append(6)
	DoublyHead.Prepend(7)
	DoublyHead.Append(8)
	DoublyHead.Prepend(9)
	DoublyHead.Append(10)
	DoublyHead.Prepend(1)

	DoublyHead.List()
	DoublyHead.Delete(1)
	DoublyHead.Delete(6)
	DoublyHead.List()
	fmt.Println("Length: ", DoublyHead.Length())
	DoublyHead.Clear()
}
