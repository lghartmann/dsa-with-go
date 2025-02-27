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

func (n *LinkedListNode) Prepend(val int) *LinkedListNode {
	newNode := &LinkedListNode{
		value: val,
		next:  n,
	}
	length++

	fmt.Println("New head: ", val)
	// we have to return the new node so we can reassign head's value to the new node
	return newNode
}

func (n *LinkedListNode) Append(val int) {
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
	// no return needed since we didn't prepend the node and the head remains the same
}

func (n *LinkedListNode) Delete(val int) *LinkedListNode {
	if n.Length() <= 1 {
		// no node will be deleted if we have a single node in our list
		fmt.Println("Your linked list has less than 2 nodes")
		return n
	}

	var oldNode *LinkedListNode = nil
	currentNode := n

	for currentNode != nil {
		if currentNode.value == val {
			if oldNode == nil {
				// since we're working with a garbage collected language we just return the next node and let it rip
				return currentNode.next
			} else {
				oldNode.next = currentNode.next
			}
			return oldNode
		}
		oldNode = currentNode
		currentNode = currentNode.next
	}
	// if node value is not found return the head calling it
	return n
}

func (n *LinkedListNode) List() {
	current := n
	for current != nil {
		fmt.Println("Value: ", current.value, "Next: ", current.next)
		current = current.next
	}
}

func (n *LinkedListNode) Clear() *LinkedListNode {
	fmt.Println("LOL NO NEED TO FREE MEMORY")
	return nil
}

func DemonstrateLinkedList() {
	Head = Head.Prepend(5)
	Head.Append(6)
	Head = Head.Prepend(7)
	Head.Append(8)
	Head = Head.Prepend(9)
	Head.Append(10)
	Head = Head.Prepend(1)

	Head = Head.Delete(1)

	Head.List()
	fmt.Println(Head.Length())
	Head = Head.Clear()
	fmt.Println("Cleared List: ", Head)
}
