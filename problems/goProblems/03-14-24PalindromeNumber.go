package goProblems

func (h *Handler) palindromeNumber(x int) (string, bool) {
	// negatives ain't palindromes
	if x < 0 {
		return "false", false
	}

	num, reversed := x, 0
	for num != 0 {
		// reverse the number
		// shift the digit by * 10
		// pop the next digit by % 10
		// join by addition
		reversed = (10 * reversed) + (num % 10)
		// shift num down by division until we run out of digits (num == 0)
		num = num / 10
	}
	// if the number equals its reverse, its a palindrome
	toReturn := (x == reversed)

	return "", toReturn
}
