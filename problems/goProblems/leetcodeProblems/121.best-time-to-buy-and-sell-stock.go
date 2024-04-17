package leetcodeProblems

/*
	@lc app=leetcode id=121 lang=golang

	[121] Best Time to Buy and Sell Stock

	https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

	algorithms
	Easy (53.61%)
	Likes:    30312
	Dislikes: 1110
	Total Accepted:    4.5M
	Total Submissions: 8.4M
	Testcase Example:  '[7,1,5,3,6,4]'

	You are given an array prices where prices[i] is the price of a given stock
	on the i^th day.

	You want to maximize your profit by choosing a single day to buy one stock
	and choosing a different day in the future to sell that stock.

	Return the maximum profit you can achieve from this transaction. If you
	cannot achieve any profit, return 0.

	Example 1:
		Input: prices = [7,1,5,3,6,4]
		Output: 5
		Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
		= 6-1 = 5.
		Note that buying on day 2 and selling on day 1 is not allowed because you
		must buy before you sell.
	Example 2:
	 	Input: prices = [7,6,4,3,1]
	 	Output: 0
	 	Explanation: In this case, no transactions are done and the max profit =
	 	0.
*/

/*
Accepted

	212/212 cases passed (79 ms)
	Your runtime beats 97.26 % of golang submissions
	Your memory usage beats 64.95 % of golang submissions (7.9 MB)
*/
func (h *LCHandler) MaxProfit(prices []int) int {
	return maxProfitPart1(prices)
}

// @lc code=start
// NOTE this problem has a followup, will need to rename
// to run uncomment below line and comment line 55
// func maxProfit(prices []int) int {
func maxProfitPart1(prices []int) int {
	// traversing the array backwards keeps it at O(n)
	currentHighest := 0
	maxProfit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		currentElement := prices[i]
		if currentElement > currentHighest {
			currentHighest = currentElement
		} else {
			currentProfit := currentHighest - currentElement
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}

// @lc code=end
