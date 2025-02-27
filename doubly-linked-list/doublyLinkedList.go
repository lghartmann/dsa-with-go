package doublyLinkedList

import "fmt"

var length = 0
var Head *DoublyLinkedList

type DoublyLinkedList struct {
	prev  *DoublyLinkedList
	value int
	next  *DoublyLinkedList
}

func (n *DoublyLinkedList) Length() int {
	return length
}

func (n *DoublyLinkedList) Prepend(val int) *DoublyLinkedList {
	current := n

	// we have to guarantee the iteration through the start of the doubly linked list
	for current.prev != nil {
		current = current.prev
	}

	current.prev = &DoublyLinkedList{
		prev:  nil,
		value: val,
		next:  current,
	}
	fmt.Println("New head: ", current.prev.value)
	return current.prev
}

func (n *DoublyLinkedList) Append(val int) {
	current := n

	// we have to guarantee the iteration through the end of the doubly linked list
	for current.next != nil {
		current = current.next
	}

	current.next = &DoublyLinkedList{
		prev:  current,
		value: val,
		next:  nil,
	}

	fmt.Println("Head remains: ", n.value)
}

func (n *DoublyLinkedList) Delete(val int) *DoublyLinkedList {
	if n.Length() <= 1 {
		// no node will be deleted if we have a single node in our list
		fmt.Println("Your linked list has less than 2 nodes")
		return n
	}

	current := n

	// we have to guarantee iteration starting from the first node
	for current.prev != nil {
		current = current.prev
	}

	for current != nil {
		if current.value == val {
			// we need to update the previous value of the next node to nil (it will be the next head)
			current.next.prev = nil
			// return the next node to reassing the head
			return current.next
		}
		current = current.next
	}

	// if node val is not found, return head
	return current
}
