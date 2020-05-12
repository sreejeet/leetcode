package main

// MinStack is a minstrack.
type MinStack struct {
	s        []int
	mins     []uint32
	count    uint32
	minCount uint32
}

// Constructor is used to initialize the data structure.
func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	ms.s = append(ms.s, x)
	ms.count++
	if ms.minCount == 0 || x < ms.s[ms.mins[ms.minCount-1]] {
		ms.mins = append(ms.mins, ms.count-1)
		ms.minCount++
	}
}

func (ms *MinStack) Pop() {
	if ms.count < 2 {
		ms.s = []int{}
		ms.mins = []uint32{}
		ms.count = 0
		ms.minCount = 0
	} else {
		if ms.mins[ms.minCount-1] == ms.count-1 {
			ms.minCount--
			ms.mins = ms.mins[:ms.minCount]
		}
		ms.count--
		ms.s = ms.s[:ms.count]
	}
}

func (ms *MinStack) Top() int {
	return ms.s[ms.count-1]
}

func (ms *MinStack) GetMin() (min int) {
	return ms.s[ms.mins[ms.minCount-1]]
}

func main() {
	obj := Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)
	println(obj.Top())
	obj.Pop()
	println(obj.GetMin())
	obj.Pop()
	println(obj.GetMin())
	obj.Pop()
	obj.Push(2147483647)
	println(obj.Top())
	println(obj.GetMin())
	obj.Push(-2147483648)
	println(obj.Top())
	println(obj.GetMin())
	obj.Pop()
	println(obj.GetMin())
}

/*
["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
*/
