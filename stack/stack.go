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
