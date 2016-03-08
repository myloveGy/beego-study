package mystack

// 定义保存元素的 stack
type Stack struct {
	i    int
	data [10]int
}

// push 将元素压入栈中
func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

// Pop从栈中弹出一个元素
func (s *Stack) Pop() (ret int) {
	s.i--
	ret = s.data[s.i]
	return
}
