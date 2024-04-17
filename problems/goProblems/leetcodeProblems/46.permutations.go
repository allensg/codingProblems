package leetcodeProblems

/*
	@lc app=leetcode id=46 lang=golang

	[46] Permutations

	https://leetcode.com/problems/permutations/description/

	algorithms
	Medium (78.25%)
	Likes:    18732
	Dislikes: 319
	Total Accepted:    2M
	Total Submissions: 2.6M
	Testcase Example:  '[1,2,3]'

	Given an array nums of distinct integers, return all the possible
	permutations. You can return the answer in any order.

	Example 1:
		Input: nums = [1,2,3]
		Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	Example 2:
		Input: nums = [0,1]
		Output: [[0,1],[1,0]]
	Example 3:
		Input: nums = [1]
		Output: [[1]]
*/

/*
Accepted
26/26 cases passed (2 ms)
Your runtime beats 74.34 % of golang submissions
Your memory usage beats 59.19 % of golang submissions (2.9 MB)
*/

func (h *LCHandler) BacktrackPermutations(nums []int) [][]int {
	return backtrackPermutations(nums)
}

// @lc code=start
// NOTE this problem has too generic a name, will need to renamed
// to run uncomment below line and comment line 42
// func permute(nums []int) [][]int {
func backtrackPermutations(nums []int) [][]int {
	results := [][]int{}
	visited := make([]bool, len(nums))
	currentPermutation := make([]int, 0, len(nums))

	var backtrack func()
	backtrack = func() {
		// basecase
		// Base Case: The base case of the recursion would be when the current permutation's
		// length equals the length of the input array nums. At this point, you add the permutation to your result set.
		if len(nums) == len(currentPermutation) {
			toAdd := make([]int, len(nums))
			copy(toAdd, currentPermutation)
			results = append(results, toAdd)
			return
		}

		//loop
		// Generate Candidates: For each position in the permutation, iterate over the elements of the input array nums, adding each element
		//that hasn't been used yet to the current permutation.
		for index, num := range nums {
			// Pruning: Similar to the combination problem, you can avoid generating permutations that cannot lead to valid solutions.
			// You can do this by maintaining a boolean array indicating which elements of nums have been used in the current permutation.
			if visited[index] {
				continue
			}
			currentPermutation = append(currentPermutation, num)
			visited[index] = true

			// Recursive Call: After adding an element to the permutation, make a recursive call to continue building the permutation.
			backtrack()
			currentPermutation = currentPermutation[:len(currentPermutation)-1]
			visited[index] = false
		}
	}

	backtrack()
	return results
}

// @lc code=end
