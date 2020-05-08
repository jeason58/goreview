package stack

type MinStack struct {
	tab    []int
	top    int
	minVal int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		tab:    []int{},
		top:    -1,
		minVal: -1,
	}
}

func (this *MinStack) Push(x int) {
	this.top++
	if this.top < cap(this.tab) {
		this.tab[this.top] = x
	} else {
		this.tab = append(this.tab, x)
	}

	if this.top == 0 {
		this.minVal = this.tab[this.top]
	} else if this.tab[this.top] < this.minVal {
		this.minVal = this.tab[this.top]
	}
}

func (this *MinStack) Pop() {
	if this.top == -1 {
		return
	}

	oldMin := this.tab[this.top]
	this.top--

	if this.top == -1 {
		this.minVal = -1
	} else if oldMin == this.minVal {
		this.minVal = this.tab[0]
		for i := 1; i <= this.top; i++ {
			if this.tab[i] < this.minVal {
				this.minVal = this.tab[i]
			}
		}
	}
}

func (this *MinStack) Top() int {
	if this.top == -1 {
		return -1
	}
	return this.tab[this.top]
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
