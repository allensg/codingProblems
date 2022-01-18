package problems

import (
	"fmt"

	"github.com/allensg/codingProblems/helpers"
)

// You are given an array of k sorted singly linked int lists. Merge the linked lists into a single sorted linked list and return it.
// Your solution should run in linear time.

// a = Node(1, Node(3, Node(5)))
// b = Node(2, Node(4, Node(6)))
// print merge([a, b])
// # 123456

func (h *Handler) SortLinkedLists(input []helpers.LinkedList) (returnString string, list helpers.LinkedList) {
	helper := helpers.Helpers{}
	returnList := helper.CreateLinkedList()

	currentNodes := make([]helpers.Node, len(input))

	// calculate length of input (total number of nodes)
	targetLen := 0
	// get the current value for each index of the array
	for i, k := range input {
		currentNodes[i] = k.GetHead()
		targetLen = targetLen + k.Length()
	}

	//while we have not sorted all the elements
	for returnList.Length() != targetLen {

		// find min of the current values
		for range currentNodes {

			minIndex := helper.FindMinIntNodeArrIndex(currentNodes)
			// add that node to the return list
			nodeToAdd := currentNodes[minIndex]
			val := nodeToAdd.GetKeyInt()
			fmt.Println("===========================")
			fmt.Println(val)

			returnList.Insert(val)
			// update currentNodes selected index with .next
			nodeToAdd.StepForward()
		}
	}

	return returnString, list
}
