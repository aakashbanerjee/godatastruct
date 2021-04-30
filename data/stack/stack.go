package stack

import (
	"fmt"
	"sync"
)

//Stack implements a LIFO data structure
type Stack struct {
	lock sync.Mutex
	i    []interface{} //could be of any type
	top  int
}

//Push inserts an element to the stack
func (s *Stack) Push(d interface{}) {
	s.lock.Lock()
	s.i = append(s.i, d)
	s.top = len(s.i)
	s.lock.Unlock()
}

//Display Stack
func (s *Stack) Display() ([]interface{}, int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.i, s.top
}

//Pop removes the last inserted element from the stack
func (s *Stack) Pop() interface{} {
	if !s.IsEmpty() {
		s.lock.Lock()
		data := s.i[s.top-1]
		s.top = len(s.i) - 1
		s.i = s.i[:s.top]
		defer s.lock.Unlock()
		return data
	}
	fmt.Println("Stack is Empty, no elements to Pop")
	return nil

}

//IsEmpty cheecks if the stack is empty
func (s *Stack) IsEmpty() bool {
	s.lock.Lock()

	if s.top <= 0 {
		defer s.lock.Unlock()
		return true
	}
	defer s.lock.Unlock()
	return false
}
