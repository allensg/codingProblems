package leetcodeProblems

/*
	@lc app=leetcode id=138 lang=golang

	[138] Copy List with Random Pointer

	https://leetcode.com/problems/copy-list-with-random-pointer/description/

	algorithms
	Medium (55.97%)
	Likes:    13600
	Dislikes: 1438
	Total Accepted:    1.2M
	Total Submissions: 2.2M
	Testcase Example:  '[[7,null],[13,0],[11,4],[10,2],[1,0]]'

	A linked list of length n is given such that each node contains an
	additional random pointer, which could point to any node in the list, or
	null.

	Construct a deep copy of the list. The deep copy should consist of exactly n
	brand new nodes, where each new node has its value set to the value of its
	corresponding original node. Both the next and random pointer of the new
	nodes should point to new nodes in the copied list such that the pointers in
	the original list and copied list represent the same list state. None of the
	pointers in the new list should point to nodes in the original list.

	For example, if there are two nodes X and Y in the original list, where
	X.random --> Y, then for the corresponding two nodes x and y in the copied
	list, x.random --> y.

	Return the head of the copied linked list.

	The linked list is represented in the input/output as a list of n nodes.
	Each node is represented as a pair of [val, random_index] where:

	val: an integer representing Node.val
	random_index: the index of the node (range from 0 to n-1) that the random
	pointer points to, or null if it does not point to any node.

	Your code will only be given the head of the original linked list.

	Example 1:
		Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
		Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	Example 2:
		Input: head = [[1,1],[2,1]]
		Output: [[1,1],[2,1]]
	Example 3:
		Input: head = [[3,null],[3,0],[3,null]]
		Output: [[3,null],[3,0],[3,null]]
*/
/*
Accepted
19/19 cases passed (0 ms)
Your runtime beats 100 % of golang submissions
Your memory usage beats 83.9 % of golang submissions (3.7 MB)
*/

func (h *LCHandler) CopyRandomList(head *Node) *Node {
	return copyRandomList(head)
}

// @lc code=start
func copyRandomList(head *Node) *Node {

	if head == nil {
		return head
	}

	// Interweave
	ptr1 := head
	for ptr1 != nil {
		newNode := &Node{
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

	return newHead
}

// @lc code=end
