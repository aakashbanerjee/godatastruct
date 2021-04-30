package stack

//Stack implements a LIFO data structure
type Stack struct {
	i   []int
	top int
}

//Push inserts an element to the stack
func (s *Stack) Push(i int) {
	s.i = append(s.i, i)
	s.top++
}

//Display Stack
func (s *Stack) Display() ([]int, int) {
	return s.i, s.top
}

//Pop removes the last inserted element from thee stack
func (s *Stack) Pop() {
	s.top--
	s.i = s.i[:s.top]
}

//IsEmpty cheecks if the stack is empty
func (s *Stack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}
