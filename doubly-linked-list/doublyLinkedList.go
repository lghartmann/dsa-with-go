package doublyLinkedList

import "fmt"

var length = 0
var Head *DoublyLinkedListNode

type DoublyLinkedListNode struct {
	prev  *DoublyLinkedListNode
	value int
	next  *DoublyLinkedListNode
}

func (n *DoublyLinkedListNode) Length() int {
	return length
}

func (n *DoublyLinkedListNode) Prepend(val int) {
	if n == nil {
		Head = &DoublyLinkedListNode{
			prev:  nil,
			value: val,
			next:  nil,
		}
		return
	}

	Head = &DoublyLinkedListNode{
		prev:  nil,
		value: val,
		next:  n,
	}

	// fmt.Println("New head: ", Head.value)
	length++
}

func (n *DoublyLinkedListNode) Append(val int) {
	if n == nil {
		Head = &DoublyLinkedListNode{
			prev:  nil,
			value: val,
			next:  nil,
		}
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

	length++
	fmt.Println("Head remains: ", n.value)
}

func (n *DoublyLinkedListNode) Delete(val int) {
	if n.Length() <= 1 {
		// no node will be deleted if we have a single node in our list
		fmt.Println("Your linked list has less than 2 nodes")
		return
	}

	current := n

	for current != nil {
		if current.value == val && current.prev == nil {
			current.next.prev = nil
			Head = current.next
		} else if current.value == val && current.next == nil {
			current.prev.next = nil
			return
		} else if current.value == val {
			current.next.prev = current.prev
			current.prev.next = current.next
			return
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
	Head = nil
}

func DemonstrateDoublyLinkedList() {
	// Head.Prepend(5)
	Head.Append(6)
	// Head.Prepend(7)
	Head.Append(8)
	// Head.Prepend(9)
	// Head.Append(10)
	// Head.Prepend(1)

	Head.List()
	// Head.Delete(1)
	// Head.Delete(6)
	// Head.List()
	fmt.Println("Length: ", Head.Length())
	Head.Clear()
}
