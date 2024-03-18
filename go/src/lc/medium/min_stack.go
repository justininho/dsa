package medium

// (155 | Medium) MinStack
// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
type MinStack struct {
	head *node
}

type node struct {
	value int
	min   int
	next  *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil || val < this.head.min {
		this.head = &node{val, val, this.head}
	} else {
		this.head = &node{val, this.head.min, this.head}
	}
}

func (this *MinStack) Pop() {
	if this.head != nil {
		this.head = this.head.next
	}
}

func (this *MinStack) Top() int {
	if this.head != nil {
		return this.head.value
	}
	var empty int
	return empty
}

func (this *MinStack) GetMin() int {
	if this.head != nil {
		return this.head.min
	}
	var empty int
	return empty
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
