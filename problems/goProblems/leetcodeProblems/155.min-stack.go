package leetcodeProblems

/*
    @lc app=leetcode id=155 lang=golang

    [155] Min Stack

    https://leetcode.com/problems/min-stack/description/

    algorithms
    Medium (54.07%)
    Likes:    13828
    Dislikes: 858
    Total Accepted:    1.6M
    Total Submissions: 3M
    Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'

    Design a stack that supports push, pop, top, and retrieving the minimum
    element in constant time.

    Implement the MinStack class:
    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

    You must implement a solution with O(1) time complexity for each
    function.

    Example 1:
        Input
        ["MinStack","push","push","push","getMin","pop","top","getMin"]
        [[],[-2],[0],[-3],[],[],[],[]]
        Output
        [null,null,null,null,-3,null,0,-2]
        Explanation
        MinStack minStack = new MinStack();
        minStack.push(-2);
        minStack.push(0);
        minStack.push(-3);
        minStack.getMin(); // return -3
        minStack.pop();
        minStack.top();    // return 0
        minStack.getMin(); // return -2
*/
/*
Accepted
31/31 cases passed (43 ms)
Your runtime beats 7.52 % of golang submissions
Your memory usage beats 84.7 % of golang submissions (7.1 MB)
*/
func (h *LCHandler) MinStackConstructor() MinStack {
	return Constructor()
}

// @lc code=start
type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	stack := []int{}
	return MinStack{stack: stack}
}

func (this *MinStack) Push(val int) {
	stack := this.stack
	this.stack = append(stack, val)
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}
	stack := this.stack
	index := len(stack) - 1
	stack = stack[:index]
	this.stack = stack
}

func (this *MinStack) Top() int {
	stack := this.stack
	return stack[len(stack)-1]
}

func (this *MinStack) GetMin() int {
	var min int
	for index, current := range this.stack {
		if index == 0 || current < min {
			min = current
		}
	}
	return min
}

func (this *MinStack) IsEmpty() bool {
	return len(this.stack) == 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
