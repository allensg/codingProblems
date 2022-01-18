package problems

import (
	"github.com/allensg/codingProblems/helpers"
)

// You are given an array of k sorted singly linked int lists. Merge the linked lists into a single sorted linked list and return it.
// Your solution should run in linear time.

// a = Node(1, Node(3, Node(5)))
// b = Node(2, Node(4, Node(6)))
// print merge([a, b])
// # 123456
// fuck link lists. just using a library next time
func (h *Handler) SortLinkedLists(input []helpers.LinkedList) (returnString string, returnList helpers.LinkedList) {
	helper := helpers.Helpers{}
	returnList = helper.CreateLinkedList()
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

			returnList.Insert(val)
			// if we have just added the last node in a sub list to the return list
			if nodeToAdd.GetNext() == nil {
				// invalidate that index
				nilNode := helper.CreateNode()
				nilNode.SetKeyInt(-1)
				currentNodes[minIndex] = nilNode
			} else {
				// update currentNodes selected index with .next
				nextNode := nodeToAdd.StepForward()
				currentNodes[minIndex] = *nextNode
			}

		}
	}

	return returnString, returnList
}
