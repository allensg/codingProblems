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

func (h *Helpers) CreateNode() (node Node) {
	return Node{}
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

// return the min value from the array of ints
func (h *Helpers) FindMinIntArrValue(input []int) (min int) {
	min = 0
	for _, k := range input {
		if k > min {
			min = k
		}
	}

	return min
}

// return the index of the min value
func (h *Helpers) FindMinIntArrIndex(input []int) int {
	minIndex, min := 0, 0
	for i, k := range input {
		if k > min {
			minIndex, min = i, k
		}
	}

	return minIndex
}

// return the min value from the array of int Nodes
func (h *Helpers) FindMinIntNodeArrValue(input []Node) int {
	min := ^int(0)
	for _, k := range input {
		key, _ := k.key.(int)
		if min > key {
			min = key
		}
	}

	return min
}

// return the index of the min value from the array of int Nodes
func (h *Helpers) FindMinIntNodeArrIndex(input []Node) int {
	minIndex, min := 0, ^int(0)
	for i, k := range input {
		key, _ := k.key.(int)
		if min > key {
			minIndex, min = i, key
		}
	}

	return minIndex
}
