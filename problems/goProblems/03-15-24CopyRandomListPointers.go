package goProblems

// A linked list of length n is given such that each node contains an additional random pointer,
// which could point to any node in the list, or null.

// Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes,
// where each new node has its value set to the value of its corresponding original node. Both the next and
// random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original
// list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

// For example, if there are two nodes X and Y in the original list, where X.random --> Y,
// then for the corresponding two nodes x and y in the copied list, x.random --> y.

// Return the head of the copied linked list.

// The linked list is represented in the input/output as a list of n nodes.
// Each node is represented as a pair of [val, random_index] where:

// val: an integer representing Node.val
// random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
// Your code will only be given the head of the original linked list.

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func (h *Handler) copyRandomList(head *RandomListNode) (string, *RandomListNode) {

	if head == nil {
		return "", head
	}

	// Interweave
	ptr1 := head
	for ptr1 != nil {
		newNode := &RandomListNode{
			Val:    ptr1.Val,
			Next:   ptr1.Next,
			Random: ptr1.Random,
		}
		ptr1.Next = newNode
		ptr1 = newNode.Next
	}

	// Update the random pointer.
	ptr1 = head
	for ptr1 != nil {
		ptr2 := ptr1.Next
		if ptr1.Random == nil {
			ptr2.Random = nil
		} else {
			ptr2.Random = ptr1.Random.Next
		}
		ptr1 = ptr1.Next.Next
	}

	// Separate the old list with new one
	ptr1 = head
	newHead := head.Next
	for ptr1 != nil {
		ptr2 := ptr1.Next
		ptr1.Next = ptr2.Next
		if ptr1.Next != nil {
			ptr2.Next = ptr1.Next.Next
		} else {
			ptr2.Next = nil
		}
		ptr1 = ptr1.Next
	}

	return "", newHead
}
