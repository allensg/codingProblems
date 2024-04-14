package goProblems

import (
	"strings"

	"github.com/allensg/codingProblems/helpers"
)

// Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system, convert it to the simplified canonical path.

// In a Unix-style file system, a period '.' refers to the current directory, a double period '..' refers to the directory up a level, and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. For this problem, any other format of periods such as '...' are treated as file/directory names.

// The canonical path should have the following format:

// The path starts with a single slash '/'.
// Any two directories are separated by a single slash '/'.
// The path does not end with a trailing '/'.
// The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
// Return the simplified canonical path.

func (h *Handler) SimplifyPath(str string) (returnString string, output string) {
	var stack helpers.StringStack

	split := strings.Split(str, "/")

	for _, val := range split {

		switch val {
		// ignore
		case "", ".":
			continue
		// everytime we hit a .. we pop a directory off the stack
		case "..":
			if !stack.IsEmpty() {
				stack.Pop()
			}
		// every time we hit a / we add that directory to the stack
		default:
			stack.Push(val)
		}

	}
	// actually because the stack is just an array we can just join strings.join /
	toReturn := strings.Join(stack, "/")
	return "", toReturn
}
