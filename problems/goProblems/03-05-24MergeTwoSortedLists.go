package goProblems

import "github.com/allensg/codingProblems/helpers"

// Definition for singly-linked list.

// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: list1 = [], list2 = []
// Output: []
// Example 3:

// Input: list1 = [], list2 = [0]
// Output: [0]
func (h *Handler) mergeTwoLists(list1 *helpers.ListNode, list2 *helpers.ListNode) string {
	mergedList := mergeList(list1, list2)
	str := generateListNodeString(mergedList)
	return str
}

// ok so apparently for the context of this problem the lists can be of different lengths between 0 and 50
// so instead of iterating over one list of the other i'll be iterating 0-50 until a both nodes are nil
func mergeList(list1 *helpers.ListNode, list2 *helpers.ListNode) *helpers.ListNode {
	baseNode := &helpers.ListNode{}
	baseNodeHead := baseNode
	node1, node2 := list1, list2
	for node1 != nil && node2 != nil {

		if node1.Val > node2.Val {
			baseNode.Val = node1.Val
			baseNode.Next = &helpers.ListNode{}
			node1 = node1.Next
		} else {
			baseNode.Val = node2.Val
			baseNode.Next = &helpers.ListNode{}
			node2 = node2.Next
		}
		baseNode = baseNode.Next
	}

	for node1 != nil {
		baseNode.Val = node1.Val
		baseNode.Next = &helpers.ListNode{}
		node1 = node1.Next
		baseNode = baseNode.Next
	}

	for node2 != nil {
		baseNode.Val = node2.Val
		baseNode.Next = &helpers.ListNode{}
		node2 = node2.Next
		baseNode = baseNode.Next
	}

	return baseNodeHead
}

func generateListNodeString(list *helpers.ListNode) string {
	valString := ""
	for list.Next != nil {
		valString = valString + string(list.Val)
	}

	return valString
}
