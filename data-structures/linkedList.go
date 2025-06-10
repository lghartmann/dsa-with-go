package datastructures

import "fmt"

var linkedLength int = 0
var LinkedHead *LinkedListNode

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func (n *LinkedListNode) Length() int {
	return linkedLength
}

func (n *LinkedListNode) Prepend(val int) {
	if n == nil {
		LinkedHead = &LinkedListNode{
			value: val,
			next:  nil,
		}
		linkedLength++
		return
	}

	newNode := &LinkedListNode{
		value: val,
		next:  n,
	}
	linkedLength++

	fmt.Println("New head: ", val)
	LinkedHead = newNode
}

func (n *LinkedListNode) Append(val int) {
	if n == nil {
		LinkedHead = &LinkedListNode{
			value: val,
			next:  nil,
		}
		linkedLength++
		return
	}

	newNode := &LinkedListNode{
		value: val,
		next:  nil,
	}
	current := n

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	fmt.Println("LinkedHead remains: ", n.value)
	linkedLength++
}

func (n *LinkedListNode) Delete(val int) {
	if n.Length() <= 1 {
		// no node will be deleted if we have a single node in our list
		fmt.Println("Your linked list has less than 2 nodes")
		return
	}

	var oldNode *LinkedListNode = nil
	currentNode := n

	for currentNode != nil {
		if currentNode.value == val {
			if oldNode == nil {
				LinkedHead = n.next
			} else {
				oldNode.next = currentNode.next
			}
			linkedLength--
		}
		oldNode = currentNode
		currentNode = currentNode.next
	}
}

func (n *LinkedListNode) List() {
	if n.Length() == 0 {
		fmt.Println("No nodes in your list")
		return
	}

	current := n
	for current != nil {
		fmt.Println("Value: ", current.value, "Next: ", current.next)
		current = current.next
	}
}

func (n *LinkedListNode) Clear() {
	fmt.Println("LOL NO NEED TO FREE MEMORY")
	LinkedHead = nil
}

func DemonstrateLinkedList() {
	LinkedHead.Prepend(5)
	LinkedHead.Append(6)
	LinkedHead.Prepend(7)
	LinkedHead.Append(8)
	LinkedHead.Prepend(9)
	LinkedHead.Append(10)
	LinkedHead.Prepend(1)

	LinkedHead.List()
	LinkedHead.Delete(1)
	LinkedHead.Delete(6)
	LinkedHead.List()
	fmt.Println(LinkedHead.Length())
	LinkedHead.Clear()
	fmt.Println("Cleared List: ", LinkedHead)
}
