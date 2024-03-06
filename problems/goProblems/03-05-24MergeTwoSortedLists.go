package goProblems

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

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
func (h *Handler) mergeTwoLists(list1 *ListNode, list2 *ListNode) string {
	mergedList := mergeList()
	str := generateListNodeString(mergedList)
	return str
}

// ok so apparently for the context of this problem the lists can be of different lengths between 0 and 50
// so instead of iterating over one list of the other i'll be iterating 0-50 until a both nodes are nil
func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	baseNode := &ListNode{}
	for list1 != nil && list2 != nil {

		if list1.Val >= list2.Val {
			baseNode.Val = list1.Val
			baseNode.Next = &ListNode{}
			list1 = list1.Next
		} else {
			baseNode.Val = list2.Val
			baseNode.Next = &ListNode{}
			list2 = list2.Next
		}
		baseNode = baseNode.Next
	}

	for list1 != nil {
		baseNode.Val = list1.Val
		baseNode.Next = &ListNode{}
		list1 = list1.Next
		baseNode = baseNode.Next
	}

	for list2 != nil {
		baseNode.Val = list2.Val
		baseNode.Next = &ListNode{}
		list2 = list2.Next
		baseNode = baseNode.Next
	}

	return baseNode
}

func generateListNodeString(list *ListNode) string {
	valString := ""
	for list.Next != nil {
		valString = valString + string(list.Val)
	}

	return valString
}
