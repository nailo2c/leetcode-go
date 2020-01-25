package problem0155

// MinStack is a struct
type MinStack struct {
	stack []item
}

type item struct {
	min int
	x   int
}

// Constructor initializes
func Constructor() MinStack {
	return MinStack{}
}

// Push add element to stack
func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}

	this.stack = append(this.stack, item{min: min, x: x})
}

// Pop drop the latest element
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

// Top return latest element
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x
}

// GetMin return minimum of the stack
func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
