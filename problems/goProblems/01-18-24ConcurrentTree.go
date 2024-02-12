package goProblems

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// https://go.dev/tour/concurrency/8

// example tree 1, 1, 2, 3, 5, 8, 13

// 		     3 		 |         8
//        /    \     |       /   \
// 	     1      8    |      3     13
//     /  \    / \   |    /  \
//    1    2  5  13  |   1    5
//                   |  / \
//                   | 1   2

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // <- closes the channel when this function returns
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) string {
	// set up channels and tres
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	tree1, tree2 := tree.New(10), tree.New(10)

	// walk both trees and write their outputs to the respective channels
	go Walk(tree1, ch1)
	go Walk(tree2, ch2)
	for val1 := range ch1 {
		select {
		case val2 := <-ch2:
			if val1 != val2 {
				return "false"
			}
		default:
			break
		}
	}
	return "true"
}

func (h *Handler) ConcurrentTree(tree1Size int, tree2Size int) (returnString string) {
	return fmt.Sprintf("10 and 10 same: %s", Same(tree.New(tree1Size), tree.New(tree2Size)))
}
