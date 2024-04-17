package leetcodeProblems

import (
	"fmt"
)

func (h *LCHandler) SummaryRanges(nums []int) []string {
	return summaryRanges(nums)
}

/*
You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).
Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b

Example 1:
	Input: nums = [0,1,2,4,5,7]
	Output: ["0->2","4->5","7"]
	Explanation: The ranges are:
	[0,2] --> "0->2"
	[4,5] --> "4->5"
	[7,7] --> "7"
Example 2:
	Input: nums = [0,2,3,4,6,8,9]
	Output: ["0","2->4","6","8->9"]
	Explanation: The ranges are:
	[0,0] --> "0"
	[2,4] --> "2->4"
	[6,6] --> "6"
	[8,9] --> "8->9"
*/

/*
Accepted
32/32 cases passed (1 ms)
Your runtime beats 80.89 % of golang submissions
Your memory usage beats 46.07 % of golang submissions (2.2 MB)
*/

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	toReturn := []string{}

	if len(nums) == 0 {
		return toReturn
	}

	rangeStart, rangeEnd := nums[0], nums[0]
	// if theres only one element don't both with any of this and just return it
	if len(nums) == 1 {
		str := formatAdd(rangeStart, rangeEnd)
		return []string{str}
	}

	for i, k := range nums {
		// if we're at the last element in the array
		if i+1 == len(nums) {
			rangeEnd = k
			toReturn = append(toReturn, formatAdd(rangeStart, rangeEnd))

		} else {
			// check the next element in the list nums[i+1]
			// if its not equal to the next value in the sequence terminate this range section
			if nums[i+1] != k+1 {
				rangeEnd = k
				toReturn = append(toReturn, formatAdd(rangeStart, rangeEnd))
				rangeStart = nums[i+1]
			}
		}
	}

	return toReturn
}

func formatAdd(start, end int) string {
	toReturn := ""
	if start != end {
		toReturn = fmt.Sprintf("%d->%d", start, end)
	} else {
		toReturn = fmt.Sprintf("%d", start)
	}

	return toReturn
}

// @lc code=end
