package leetcodeProblems

/*
	@lc app=leetcode id=141 lang=golang

	[141] Linked List Cycle

	https://leetcode.com/problems/linked-list-cycle/description/

	algorithms
	Easy (50.34%)
	Likes:    15225
	Dislikes: 1320
	Total Accepted:    2.9M
	Total Submissions: 5.8M
	Testcase Example:  '[3,2,0,-4]\n1'

	Given head, the head of a linked list, determine if the linked list has a
	cycle in it.

	There is a cycle in a linked list if there is some node in the list that can
	be reached again by continuously following the next pointer. Internally, pos
	is used to denote the index of the node that tail's next pointer is
	connected to. Note that pos is not passed as a parameter.

	Return true if there is a cycle in the linked list. Otherwise, return
	false.

	Example 1:
		Input: head = [3,2,0,-4], pos = 1
		Output: true
		Explanation: There is a cycle in the linked list, where the tail connects to
		the 1st node (0-indexed).
	Example 2:
		Input: head = [1,2], pos = 0
		Output: true
		Explanation: There is a cycle in the linked list, where the tail connects to
		the 0th node.
	Example 3:
		Input: head = [1], pos = -1
		Output: false
		Explanation: There is no cycle in the linked list.
*/
/*
Accepted
29/29 cases passed (3 ms)
Your runtime beats 94.68 % of golang submissions
Your memory usage beats 19.69 % of golang submissions (6 MB)
*/

// @lc code=start
// NOTE this problem has too generic a name, will need to rename
// to run uncomment below line and comment line 59
func hasCycle(head *ListNode) bool {
	// func linkedListCycle(head *ListNode) bool {
	passed := make(map[*ListNode]bool)
	current := head
	if head == nil {
		return false
	}

	for current.Next != nil {
		if _, found := passed[current]; found {
			return true
		}
		passed[current] = true
		current = current.Next
	}

	// if we hit next == nil we have hit the end with no cycle.
	return false
}

// @lc code=end
