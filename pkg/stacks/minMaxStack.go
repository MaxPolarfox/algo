package stacks

type MinMaxStack struct {
	stack       []int
	minMaxStack []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

func (stack *MinMaxStack) Peek() int {
	return stack.stack[len(stack.stack)-1]
}

func (stack *MinMaxStack) Pop() int {
	popped := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	return popped
}

func (stack *MinMaxStack) Push(number int) {
	newMinMax := entry{min: number, max: number}
	if len(stack.minMaxStack) > 0 {
		oldMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(oldMinMax.min, number)
		newMinMax.max = max(oldMinMax.max, number)
	}
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
	stack.stack = append(stack.stack, number)
}

func (stack *MinMaxStack) GetMin() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
