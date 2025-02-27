package linkedList

import "fmt"

var length int = 0
var Head *LinkedListNode

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func (n *LinkedListNode) Length() int {
	return length
}

func (n *LinkedListNode) Prepend(val int) {
	if n == nil {
		Head = &LinkedListNode{
			value: val,
			next:  nil,
		}
		return
	}

	newNode := &LinkedListNode{
		value: val,
		next:  n,
	}
	length++

	fmt.Println("New head: ", val)
	Head = newNode
}

func (n *LinkedListNode) Append(val int) {
	if n == nil {
		Head = &LinkedListNode{
			value: val,
			next:  nil,
		}
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
	fmt.Println("Head remains: ", n.value)
	length++
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
				Head = n.next
			} else {
				oldNode.next = currentNode.next
			}
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
	Head = nil
}

func DemonstrateLinkedList() {
	Head.Prepend(5)
	Head.Append(6)
	Head.Prepend(7)
	Head.Append(8)
	Head.Prepend(9)
	Head.Append(10)
	Head.Prepend(1)

	Head.List()
	Head.Delete(1)
	Head.Delete(6)
	Head.List()
	fmt.Println(Head.Length())
	Head.Clear()
	fmt.Println("Cleared List: ", Head)
}
