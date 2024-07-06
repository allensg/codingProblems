package leetcodeProblems

import "sort"

/*
	@lc app=leetcode id=846 lang=golang

	[846] Hand of Straights

	https://leetcode.com/problems/hand-of-straights/description/

	algorithms
	Medium (72.57%)

	Alice has some number of cards and she wants to rearrange the cards into groups so that
	each group is of size groupSize, and consists of groupSize consecutive cards.

	Given an integer array hand where hand[i] is the value written on the ith card and
	an integer groupSize, return true if she can rearrange the cards, or false otherwise.


	Example 1:
		Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
		Output: true
		Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]

	Example 2:
		Input: hand = [1,2,3,4,5], groupSize = 4
		Output: false
		Explanation: Alice's hand can not be rearranged into groups of 4.

*/

// @lc code=start
func (h *LCHandler) IsNStraightHand(hand []int, groupSize int) bool {
	return isNStraightHand(hand, groupSize)
}

func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	// make a table
	table := [][]int{}

	for handIndex, toInsert := range hand {
		if handIndex == 0 {
			table = append(table, []int{toInsert})
			continue
		}

		appended := false
		// come across an int, check the table.
		for tableIndex, tableRow := range table {
			// check each row we've initialized so far, this will get bigger as we go.
			// grab the last element on each slice
			toCheck := tableRow[len(tableRow)-1]
			// if you can add it to the end of an one of the existing arrays sequentially
			// && it doesn't exceed the array size, add it
			compare := toCheck + 1
			if compare == toInsert && len(tableRow) < groupSize {
				tableRow = append(tableRow, toInsert)
				table[tableIndex] = tableRow
				appended = true
				break
			}
		}
		if !appended {
			// if it doesnt, add as the first element in a new array
			newRow := []int{toInsert}
			table = append(table, newRow)
		}
	}

	// if any row in the table is not equal to the group size, you can't do it.
	for _, row := range table {
		if len(row) != groupSize {
			return false
		}
	}

	return true
}

// @lc code=end
