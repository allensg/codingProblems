package goProblems

import "github.com/allensg/codingProblems/helpers"

// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.

// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
func (h *Handler) linkedListCycle(head *helpers.ListNode) (string, bool) {
	passed := make(map[*helpers.ListNode]bool)
	current := head
	if head == nil {
		return "", false
	}

	for current.Next != nil {
		if _, found := passed[current]; found {
			return "", true
		}
		passed[current] = true
		current = current.Next
	}

	// if we hit next == nil we have hit the end with no cycle.
	return "", true
}
