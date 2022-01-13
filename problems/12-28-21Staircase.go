package problems

import (
	"fmt"
)

// You are given a positive integer N which represents the number of steps in a staircase.
// You can either climb 1 or 2 steps at a time.
// Write a function that returns the number of unique ways to climb the stairs.
func (h *Handler) Staircase(inputSteps int) (returnString string, answer int) {
	// looks like this is a fibbonaci problem so its just what is the fib. sequence at N stairs
	// so for inputs: 0 | 1 | 2 | 3 | 4 | 5 | 6 |  7 |  8 |  9 | 10 | 11 |  12 |  13 |  14 |...
	// we should see: 0 | 1 | 1 | 2 | 3 | 5 | 8 | 13 | 21 | 34 | 55 | 89 | 144 | 233 | 377 |...

	// recursion would work but would be runtime O(n^2)
	// answer = StaircaseRec(inputSteps)

	// this is less clean programatically than the recursive method
	// but its a single loop iteration so its O(n)
	if inputSteps > 1 {
		// initialize starting at case 1
		first, second, sum := 0, 1, 0
		for i := 1; i < inputSteps; i++ {
			sum = first + second
			first, second = second, sum
		}
		answer = sum
	} else {
		// 2 is the divergence point so 0, 1 will all match up to the number of steps.
		answer = inputSteps
	}

	returnString = fmt.Sprintf("For a staircase of size %d, there are %d unique ways to climb it", inputSteps, answer)
	return returnString, answer
}

func StaircaseRec(inputSteps int) int {
	if inputSteps <= 1 {
		return inputSteps
	}

	return StaircaseRec(inputSteps-1) + StaircaseRec(inputSteps-2)
}
