package goProblems

// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
// You may return the answer in any order.

// Example 1:

// Input: n = 4, k = 2
// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Explanation: There are 4 choose 2 = 6 total combinations.
// Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
// Example 2:

// Input: n = 1, k = 1
// Output: [[1]]
// Explanation: There is 1 choose 1 = 1 total combination.

// algorithm suggested by gpt. looking at the results on leetcode submissions this seems to line up pretty
// well with what other people came up with (or asked gpt to come up with)

//	Use Backtracking: This problem can be efficiently solved using a backtracking algorithm. Backtracking is a technique
// 		for generating all possible solutions to a problem by building candidates incrementally and
// 		abandoning a candidate as soon as it's determined that it cannot lead to a valid solution.
// Recursive Approach: You can define a recursive function that generates combinations by trying each possible
// 		 candidate for each position in the combination.
// Base Case: The base case of the recursion would be when the length of the combination equals k.
//		 At this point, you add the combination to your result set.
// Generate Candidates: For each position in the combination, you would iterate over the numbers from 1 to n, adding each number that
// 		 hasn't been used yet to the combination.
// Pruning: To optimize the algorithm, you can avoid generating combinations that cannot lead to valid solutions.
//		 For instance, if the number of elements remaining to choose (n - currentNumber) is less than the number of
//		 positions left in the combination (k - len(currentCombination)), you can stop exploring that branch.
// Recursive Call: After adding a number to the combination, you make a recursive call to continue building the combination.
// Backtrack: After the recursive call returns, you remove the last element added
//		 to the combination, allowing you to try other candidates for that position.
// Return Result: After exploring all possible combinations, return the list of combinations as the final result.

// By following these steps, you can efficiently generate all possible combinations of k numbers chosen from the range [1, n].
//
//	This approach ensures that you avoid generating duplicate combinations and that you explore only the branches of
//	 the search tree that can lead to valid solutions.
func (h *Handler) backtrackCombinations(n int, k int) (string, [][]int) {
	// final array of results
	results := [][]int{}
	currentCombos := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// if we've hit the end of our current combination subset
		if len(currentCombos) == k {
			// add that subset to the resultset
			combo := make([]int, k)
			copy(combo, currentCombos)
			results = append(results, combo)
			return
		}

		// for each position, iterate 1-n adding unused numbers to the result set
		for i := start; n > i; i++ {
			// add to the combos array
			currentCombos = append(currentCombos, i)
			// recurse
			backtrack(i + 1)
			// move the range up 1
			currentCombos = currentCombos[:len(currentCombos)-1]
		}
	}

	backtrack(1)
	return "", results
}
