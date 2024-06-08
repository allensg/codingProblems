package leetcodeProblems

import (
	"fmt"
	"math"
)

/*
	@lc app=leetcode id=23 lang=golang

	[23] Merge K Sorted Lists

	https://leetcode.com/problems/merge-k-sorted-lists/description/

	algorithms
	Hard

	You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

	Merge all the linked-lists into one sorted linked-list and return it.
	Example 1:
		Input: lists = [[1,4,5],[1,3,4],[2,6]]
		Output: [1,1,2,3,4,4,5,6]
		Explanation: The linked-lists are:
		[
		1->4->5,
		1->3->4,
		2->6
		]
		merging them into one sorted list:
		1->1->2->3->4->4->5->6
	Example 2:
		Input: lists = []
		Output: []
	Example 3:
		Input: lists = [[]]
		Output: []
*/

func (h *LCHandler) MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

// idk almost works or something
// @lc code=start
func mergeKLists(lists []*ListNode) *ListNode {
	baseNode, processedMap := &ListNode{}, map[int]bool{}
	toReturn := baseNode // keep track of the base node head.

	valid, validationReturn := validate(lists)
	if !valid {
		return validationReturn
	}

	minHead := func(lists []*ListNode) int {
		min, minIndex, hasNodes := math.MaxInt, 0, false
		// if a list is nil, the loop will skip it
		for index, list := range lists {
			if list == nil {
				processedMap[index] = true
				continue
			}

			hasNodes = true
			if min == math.MaxInt || list.Val < min {
				min = list.Val
				minIndex = index
			}
		}
		// if we pass in an array with all of the points at nil (we have processed all elements)
		// return -1 to indicate we break out of of the loop
		if !hasNodes {
			minIndex = -1
		}
		fmt.Println("hasnodes")
		return minIndex
	}

	loopCount := 0
	for true {
		fmt.Println("==============================================")
		fmt.Printf("loopcount: %d\n", loopCount)
		loopCount += 1
		// get the index of the min value
		toAddIndex := minHead(lists)
		fmt.Printf("toAddIndex %d\n", toAddIndex)
		// if we're out of nodes to iterate over, exit the loop
		if toAddIndex == -1 {
			return toReturn
		}

		// get the node
		toAdd := lists[toAddIndex]

		baseNode.Val = toAdd.Val
		fmt.Printf("updated basenode value to: %d\n", baseNode.Val)
		// update the base node
		if len(processedMap) <= len(lists)-1 {
			fmt.Println("added basenode.next\n")
			baseNode.Next = &ListNode{}
			baseNode = baseNode.Next
		}

		// move the one we grabbed forward an index
		fmt.Printf("updated toAdd at index: %d from: %d\n", toAddIndex, toAdd.Val)
		toAdd = toAdd.Next
		lists[toAddIndex] = toAdd

	}

	return toReturn
}

func validate(lists []*ListNode) (valid bool, toReturn *ListNode) {
	if len(lists) == 0 { // is it empty
		return false, nil
	}

	if len(lists) == 1 { // is there only one
		return false, lists[0]
	}

	hasValues := false
	for _, list := range lists {
		if list == nil {
			continue
		}
		hasValues = true
	}

	if !hasValues {
		return false, nil // is it an array of empty lists for some fucking reason
	}

	return true, nil
}

// @lc code=end
