package goProblems

// You are given a 2D array of integers.
// Print out the clockwise spiral traversal of the matrix.

// Input: [
//	[1,  2,  3,  4,  5],
//	[6,  7,  8,  9,  10],
//	[11, 12, 13, 14, 15],
//	[16, 17, 18, 19, 20]]
//
// Output: 1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12

// ex Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
func (h *Handler) MatrixSpiralPrint(matrix [][]int) (returnString string, answer []int) {

	returnString = "Pythag tripples for set: "

	if len(matrix) == 0 {
		return "", []int{}
	}

	var result []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse rightmost column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Ensure we haven't crossed boundaries
		if top <= bottom {
			// Traverse bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// Traverse leftmost column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return "", result
}
