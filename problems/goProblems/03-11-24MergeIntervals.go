package goProblems

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

	
	// i think the solution for this is to check the ranges and merge. 
	// so if the end of range 1 is greater than the start of range 2, merge them
	// then repeat for each element
	rangeStart, rangeEnd := intervals[0][0], intervals[0][1]
	for i := 0; len(intervals) > i; i++ {
		// leftStart, leftEnd := intervals[]
		rightStart, rightEnd := intervals[]		

	}

	return "", toReturn
}
