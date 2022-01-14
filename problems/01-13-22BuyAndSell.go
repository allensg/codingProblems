package problems

import "fmt"

// You are given an array. Each element represents the price of a stock on that particular day.
// Calculate and return the maximum profit you can make from buying and selling that stock only once

// Input: [9, 11, 8, 5, 7, 10]
// Output: 10-5 = 5
func (h *Handler) BuyAndSell(input []int) (returnString string, maxProfit int) {

	// traversing the array backwards keeps it at O(n)
	currentHighest := 0
	maxProfit = 0
	for i := len(input) - 1; i >= 0; i-- {
		currentElement := input[i]
		if currentElement > currentHighest {
			currentHighest = currentElement
		} else {
			currentProfit := currentHighest - currentElement
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	strArr := h.IntArrToString(input)
	returnString = fmt.Sprintf("The Max Profit that can be gained on day %s is %d", strArr, maxProfit)

	return returnString, maxProfit
}
