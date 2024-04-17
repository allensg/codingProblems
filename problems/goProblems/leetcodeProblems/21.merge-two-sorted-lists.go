package leetcodeProblems

/*
	@lc app=leetcode id=21 lang=golang

	[21] Merge Two Sorted Lists

	https://leetcode.com/problems/merge-two-sorted-lists/description/

	algorithms
	Easy (64.20%)
	Likes:    21379
	Dislikes: 2061
	Total Accepted:    4.1M
	Total Submissions: 6.4M
	Testcase Example:  '[1,2,4]\n[1,3,4]'

	You are given the heads of two sorted linked lists list1 and list2.

	Merge the two lists into one sorted list. The list should be made by
	splicing together the nodes of the first two lists.

	Return the head of the merged linked list.

	Example 1:
		Input: list1 = [1,2,4], list2 = [1,3,4]
		Output: [1,1,2,3,4,4]
	Example 2:
		Input: list1 = [], list2 = []
		Output: []
	Example 3:
		Input: list1 = [], list2 = [0]
		Output: [0]
*/
/*
Accepted
208/208 cases passed (3 ms)
Your runtime beats 62.24 % of golang submissions
Your memory usage beats 7.86 % of golang submissions (2.7 MB)
*/
func (h *LCHandler) MergeTwoSinglyLinkedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return mergeTwoSinglyLinkedLists(list1, list2)
}

// @lc code=start
// NOTE this problem has too generic a name, will need to rename
// to run uncomment below line and comment line 59
func mergeTwoSinglyLinkedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	baseNode := &ListNode{}
	current := baseNode
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return baseNode.Next
}

// @lc code=end
