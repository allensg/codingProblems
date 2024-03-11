package goProblems

// You are given an array. Each element represents the price of a stock on that particular day.
// Calculate and return the maximum profit you can make from buying and selling that stock only once

// Input: [9, 11, 8, 5, 7, 10]
// Output: 10-5 = 5
// func (h *Handler) BuyAndSell(input []int) (returnString string, maxProfit int) {
// 	helper := &helpers.Helpers{}
// 	// traversing the array backwards keeps it at O(n)
// 	currentHighest := 0
// 	maxProfit = 0
// 	for i := len(input) - 1; i >= 0; i-- {
// 		currentElement := input[i]
// 		if currentElement > currentHighest {
// 			currentHighest = currentElement
// 		} else {
// 			currentProfit := currentHighest - currentElement
// 			if currentProfit > maxProfit {
// 				maxProfit = currentProfit
// 			}
// 		}
// 	}

// 	strArr := helper.IntArrToString(input)
// 	returnString = fmt.Sprintf("The Max Profit that can be gained on day %s is %d", strArr, maxProfit)

// 	return returnString, maxProfit
// }

// max profit part 2

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
// However, you can buy it then immediately sell it on the same day.

// Find and return the maximum profit you can achieve.

// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.

// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Total profit is 4.
func (h *Handler) BuyAndSell(input []int) (returnString string, maxProfit int) {
	return "", 0
}
