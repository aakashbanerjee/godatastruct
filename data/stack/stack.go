package stack

import "fmt"

//Stack implements a LIFO data structure
type Stack struct {
	i   []interface{} //could be of any type
	top int
}

//Push inserts an element to the stack
func (s *Stack) Push(d interface{}) {
	s.i = append(s.i, d)
	s.top = len(s.i)
}

//Display Stack
func (s *Stack) Display() ([]interface{}, int) {
	return s.i, s.top
}

//Pop removes the last inserted element from the stack
func (s *Stack) Pop() interface{} {
	if !s.IsEmpty() {
		data := s.i[s.top-1]
		s.top = len(s.i) - 1
		s.i = s.i[:s.top]
		return data
	}
	fmt.Println("Stack is Empty, no elements to Pop")
	return nil

}

//IsEmpty cheecks if the stack is empty
func (s *Stack) IsEmpty() bool {
	if s.top <= 0 {
		return true
	}
	return false
}
