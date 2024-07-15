package stack

type MinStack struct {
	// Head of the stack
	head *StackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	// For the very first node
	if this.head == nil {
		this.head = &StackNode{value: val, minValue: val, next: nil}
	} else {
		// For the subsequent nodes
		this.head = &StackNode{value: val, minValue: min(val, this.head.minValue), next: this.head}
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.value
}

func (this *MinStack) GetMin() int {
	return this.head.minValue
}

type StackNode struct {
	value    int
	minValue int
	next     *StackNode
}
