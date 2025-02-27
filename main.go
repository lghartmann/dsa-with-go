package main

import (
	"fmt"
	"os"
	"strings"

	doublyLinkedList "github.com/lghartmann/data-structures-in-go/doubly-linked-list"
	linkedList "github.com/lghartmann/data-structures-in-go/linked-list"
)

func main() {
	// args := os.Args

	// if len(args) != 2 {
	// 	fmt.Println("Correct usage: `go run main.go DATA_STRUCTURE_HERE`")
	// 	os.Exit(1)
	// }

	var teste = "doubly_linked_list"

	switch strings.ToLower(teste) {
	case "linked_list":
		fmt.Println("Begin linked list:")
		linkedList.DemonstrateLinkedList()
	case "try":
		fmt.Println("Try")
	case "doubly_linked_list":
		fmt.Println("Begin Doubly Linked List:")
		doublyLinkedList.DemonstrateDoublyLinkedList()
	default:
		fmt.Println("The options are: \n", "linked_list\n", "try\n", "doubly_linked_list\n")
		os.Exit(1)
	}
}
