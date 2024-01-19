package problems

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
	defer close(ch)
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// set up channels and tres
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	tree1, tree2 := tree.New(10), tree.New(10)

	// walk both trees and write their outputs to the respective channels
	go Walk(tree1, ch1)
	go Walk(tree2, ch2)

	for {
		c1, ok1 := <-ch1
		c2, ok2 := <-ch2

		if c1 != c2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func (h *Handler) ConcurrentTree() (returnString string) {
	fmt.Println("10 and 10 same: ", Same(tree.New(10), tree.New(10)))
	return ""
}
