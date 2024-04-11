package goProblems

import (
	"github.com/allensg/codingProblems/helpers"
)

func (h *Handler) IsValidParens(str string) (returnString string, isValid bool) {
	var stack helpers.Stack

	leftMap := map[rune]bool{'{': true, '[': true, '(': true}
	mapping := map[rune]rune{'}': '{', ']': '[', ')': '('}

	for _, char := range str {

		if _, foundLeft := leftMap[char]; foundLeft {
			stack.Push(char)
		}

		if expectedBracket, foundRight := mapping[char]; foundRight {

			if stack.IsEmpty() || expectedBracket != stack[len(stack)-1] {
				return "", false
			}
			stack.Pop()
		}
	}

	return "", stack.IsEmpty()
}
