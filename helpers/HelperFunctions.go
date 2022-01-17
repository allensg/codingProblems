package helpers

import (
	"fmt"
	"strings"
)

type (
	// Helpers will act as an entry point to anything i need for
	// data generation and test suite control
	Helpers struct {
	}
)

func (h *Helpers) IntArrToString(input []int) (output string) {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ","), "[]")
}

func (h *Helpers) AreIntArraysEqual(a []int, b []int) (equal bool) {
	equal = true
	if len(a) != len(b) {
		equal = false
	} else {
		for i, k := range a {
			if k != b[i] {
				equal = false
			}
		}
	}

	return equal
}

func (h *Helpers) CreateLinkedList() (list LinkedList) {
	return LinkedList{}
}

func (h *Helpers) AreLinkedListsEqual(a LinkedList, b LinkedList) (equal bool) {
	strA, strB := a.DisplayString(), b.DisplayString()
	if strA != strB {
		return false
	}
	return true
}

func (h *Helpers) IntArrToLinkedList(input []int) (list LinkedList) {
	list = LinkedList{}

	for _, k := range input {
		list.Insert(k)
	}

	return list
}

func (h *Helpers) StringArrToLinkedList(input []string) (list LinkedList) {
	list = LinkedList{}

	for _, k := range input {
		list.Insert(k)
	}

	return list
}
