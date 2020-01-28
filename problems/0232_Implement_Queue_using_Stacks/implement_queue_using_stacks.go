package problem0232

// MyQueue is a struct
type MyQueue struct {
	queue []int
}

// Constructor initialize
func Constructor() MyQueue {
	return MyQueue{}
}

// Push element x to the back of queue.
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

// Pop removes the element from in front of queue and returns that element.
func (this *MyQueue) Pop() int {
	popElement := this.queue[0]
	this.queue = this.queue[1:]
	return popElement
}

// Peek get the front element.
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

// Empty returns whether the queue is empty.
func (this *MyQueue) Empty() bool {
	if len(this.queue) > 0 {
		return false
	}

	return true
}
