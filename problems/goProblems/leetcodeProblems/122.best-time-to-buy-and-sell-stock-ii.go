package leetcodeProblems

/*
	@lc app=leetcode id=122 lang=golang

	[122] Best Time to Buy and Sell Stock II

	https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

	algorithms
	Medium (66.36%)
	Likes:    13292
	Dislikes: 2678
	Total Accepted:    1.9M
	Total Submissions: 2.8M
	Testcase Example:  '[7,1,5,3,6,4]'

	You are given an integer array prices where prices[i] is the price of a
	given stock on the i^th day.

	On each day, you may decide to buy and/or sell the stock. You can only hold
	at most one share of the stock at any time. However, you can buy it then
	immediately sell it on the same day.

	Find and return the maximum profit you can achieve.

	Example 1:
		Input: prices = [7,1,5,3,6,4]
		Output: 7
		Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
		Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
		Total profit is 4 + 3 = 7.
	Example 2:
		Input: prices = [1,2,3,4,5]
		Output: 4
		Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
		= 5-1 = 4.
		Total profit is 4.
	Example 3:
		Input: prices = [7,6,4,3,1]
		Output: 0
		Explanation: There is no way to make a positive profit, so we never buy the
		stock to achieve the maximum profit of 0.
*/
/*
Accepted
200/200 cases passed (5 ms)
Your runtime beats 47.79 % of golang submissions
Your memory usage beats 8.78 % of golang submissions (3.4 MB)
*/

func (h *LCHandler) MaxProfitPt2(prices []int) int {
	return maxProfit2(prices)
}

// @lc code=start
// NOTE this is a followup to another problem, will need to rename
// to run uncomment below line and comment line 54
// func maxProfit(prices []int) int {
func maxProfit2(prices []int) int {
	// initial thoughs for this was to determine some kind of average with preprocessing
	// then with high and low use that average to that to determine sales on a per iteration basis
	// when that didn't work i found this and it works great with a simple lookback.
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}

// @lc code=end
