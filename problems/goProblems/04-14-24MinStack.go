package goProblems

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

// most of my elements will exist in helpers.StringStack
func (h *Handler) MinStack(str string) (returnString string, output string) {
	// var stack helpers.IntStack
	// var m int
	return "", ""
}

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	stack := []int{}
	return MinStack{stack: stack}
}

// func (this *MinStack) Top() int {
//     stack := this.stack
//     return stack[len(stack)-1]
// }

// func (this *MinStack) GetMin() int {
//     var min int
//     for index, current := range this.stack {
//         if index==0 || current < min {
//             min = current
//         }
//     }
//     return min
// }
