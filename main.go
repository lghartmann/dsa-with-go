package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lghartmann/dsa-with-go/algorithms"
	datastructures "github.com/lghartmann/dsa-with-go/data-structures"
)

func main() {
	// args := os.Args

	// if len(args) != 2 {
	// 	fmt.Println("Correct usage: `go run main.go DATA_STRUCTURE_HERE`")
	// 	os.Exit(1)
	// }

	testeInput := []int{8, 6, 5, 7, 4, 9, 1, 2, 3}

	fmt.Println(testeInput)

	algorithms.MergeSort(testeInput)

	fmt.Println(testeInput)

	var teste = "doubly_linked_list"

	switch strings.ToLower(teste) {
	case "linked_list":
		fmt.Println("Begin linked list:")
		datastructures.DemonstrateLinkedList()
	case "try":
		fmt.Println("Try")
	case "doubly_linked_list":
		fmt.Println("Begin Doubly Linked List:")
		datastructures.DemonstrateDoublyLinkedList()
	default:
		fmt.Println("The options are: \n", "linked_list\n", "try\n", "doubly_linked_list\n")
		os.Exit(1)
	}
}
