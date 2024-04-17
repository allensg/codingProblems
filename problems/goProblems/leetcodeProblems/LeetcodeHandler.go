package leetcodeProblems

import (
	"github.com/labstack/echo/v4"
)

type (
	// Handler contains reference to whatever I need specifically for leetcode
	// lots of problems reference earlier problems so this is a way to tie that all
	// together
	LCHandler struct {
		Env    map[string]string
		Logger echo.Context
	}
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a Node, specifically LC138, random pointer.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type Stack []rune
type StringStack []string
type IntStack []int

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// string
func (s *StringStack) Push(val string) {
	*s = append(*s, val)
}

func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s StringStack) IsEmpty() bool {
	return len(s) == 0
}

// int
func (s *IntStack) Push(val int) {
	*s = append(*s, val)
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s IntStack) IsEmpty() bool {
	return len(s) == 0
}
