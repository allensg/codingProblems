package leetcodeProblems

import (
	"math"
)

/*
	@lc app=leetcode id=202 lang=golang

	[202] Happy Number

	https://leetcode.com/problems/happy-number/description/

	algorithms
	Easy (56.04%)
	Likes:    10126
	Dislikes: 1401
	Total Accepted:    1.4M
	Total Submissions: 2.6M
	Testcase Example:  '19'

	Write an algorithm to determine if a number n is happy.

	A happy number is a number defined by the following process:

	Starting with any positive integer, replace the number by the sum of the
	squares of its digits.
	Repeat the process until the number equals 1 (where it will stay), or it
	loops endlessly in a cycle which does not include 1.
	Those numbers for which this process ends in 1 are happy.

	Return true if n is a happy number, and false if not.

	Example 1:
		Input: n = 19
		Output: true
		Explanation:
		1^2 + 9^2 = 82
		8^2 + 2^2 = 68
		6^2 + 8^2 = 100
		1^2 + 0^2 + 0^2 = 1
	Example 2:
		Input: n = 2
		Output: false
*/
/*
Accepted
420/420 cases passed (4 ms)
Your runtime beats 34.8 % of golang submissions
Your memory usage beats 5.1 % of golang submissions (2.7 MB)
*/

func (h *LCHandler) IsHappyNumber(n int) bool {
	return isHappy(n)
}

// @lc code=start
type mathStuff struct {
	digits    []int
	squareSum int
}

func isHappy(n int) bool {
	// make a map of ints and their corresponding digit array for loop detection
	hitMap := map[int]*mathStuff{}

	currentNum := n
	for true {
		_, found := hitMap[currentNum]
		// if we find a value we've already hit we're in a loop and should exit the program
		if found {
			return false
		}

		// split the current int into its component digits and squared sum
		numMathStuff := numSplit(currentNum)
		// success case
		if numMathStuff.squareSum == 1 {
			return true
		}

		// else update the hitmap
		hitMap[currentNum] = numMathStuff
		// and the current num and continue
		currentNum = numMathStuff.squareSum

	}

	return false
}

func numSplit(n int) *mathStuff {
	digits, squareSum := []int{}, 0
	for n != 0 {
		digit := n % 10
		digits = append(digits, digit)
		// shift num down by division until we run out of digits (num == 0)
		n = n / 10

		digitSquared := math.Pow(float64(digit), 2)
		squareSum = squareSum + int(digitSquared)
	}

	return &mathStuff{
		digits:    digits,
		squareSum: squareSum,
	}
}

// @lc code=end
