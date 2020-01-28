package problem0225

// MyStack is a struct
type MyStack struct {
	stack []int
}

// Constructor initializes
func Constructor() MyStack {
	return MyStack{}
}

// Push element x onto stack.
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

// Pop Removes the element on top of the stack and returns that element.
func (this *MyStack) Pop() int {
	popElement := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return popElement
}

// Top get the top element.
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// Empty Returns whether the stack is empty.
func (this *MyStack) Empty() bool {
	if len(this.stack) > 0 {
		return false
	}

	return true
}
