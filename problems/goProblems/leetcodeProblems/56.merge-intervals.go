package leetcodeProblems

import "sort"

/*
	@lc app=leetcode id=56 lang=golang

	[56] Merge Intervals

	https://leetcode.com/problems/merge-intervals/description/

	algorithms
	Medium (47.26%)
	Likes:    21852
	Dislikes: 760
	Total Accepted:    2.3M
	Total Submissions: 4.9M
	Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'

	Given an array of intervals where intervals[i] = [starti, endi], merge all
	overlapping intervals, and return an array of the non-overlapping intervals
	that cover all the intervals in the input.

	Example 1:
		Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
		Output: [[1,6],[8,10],[15,18]]
		Explanation: Since intervals [1,3] and [2,6] overlap, merge them into
		[1,6].
	Example 2:
		Input: intervals = [[1,4],[4,5]]
		Output: [[1,5]]
		Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/
/*
Accepted
170/170 cases passed (17 ms)
Your runtime beats 58.72 % of golang submissions
Your memory usage beats 85.05 % of golang submissions (6.3 MB)
*/

func (h *LCHandler) IntervalMerge(intervals [][]int) [][]int {
	return intervalMerge(intervals)
}

// @lc code=start
// NOTE this problem has too generic a name, will need to rename
// to run uncomment below line and comment line 40
// func merge(intervals [][]int) [][]int {
func intervalMerge(intervals [][]int) [][]int {
	toReturn := [][]int{}
	if len(intervals) == 0 {
		return toReturn
	}

	// Sort the slice (array) of intervals intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	toReturn = append(toReturn, intervals[0])

	// We iterate on every items (except the first one) of the intervals slice
	// and we compare it with the last item of the mergedIntervals slice
	for _, interval := range intervals[1:] {
		mergedEnd := toReturn[len(toReturn)-1]

		if interval[0] > mergedEnd[1] {
			// If the interval "completely overlap" append interval, overwriting mergedEnd on the next iteration
			toReturn = append(toReturn, interval)
		} else if interval[1] > mergedEnd[1] {
			// Else if interval "overlap" (but not completely) mergeEnd: change the 1 value of interval for the 1 value of top
			mergedEnd[1] = interval[1]
		}
	}
	return toReturn
}

// @lc code=end
