package problems

import (
	"github.com/allensg/codingProblems/helpers"
)

// You are given an array of k sorted singly linked lists. Merge the linked lists into a single sorted linked list and return it.
// Your solution should run in linear time.

// a = Node(1, Node(3, Node(5)))
// b = Node(2, Node(4, Node(6)))
// print merge([a, b])
// # 123456

func (h *Handler) SortLinkedLists(input []helpers.LinkedList) (returnString string, list helpers.LinkedList) {
	// helpers := helpers.Helpers{}
	for _, value := range input {
		helpers.ShowBackwards(value)
	}
	return returnString, list
}
