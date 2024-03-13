package goProblems

import "sort"

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
func (h *Handler) mergeInterval(intervals [][]int) (string, [][]int) {
	toReturn := [][]int{}
	if len(intervals) == 0 {
		return "", toReturn
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
	return "", toReturn
}
