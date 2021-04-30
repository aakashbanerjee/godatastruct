package main

import (
	"fmt"

	"github.com/godatastruct/data/stack"
)

func main() {
	callStack() //Executes all func calls for Stack
}

func printStack(st []interface{}, t int) {
	fmt.Printf("Stack Elements: %v Number of elements: %v\n", st, t)
}

func callStack() {
	s := &stack.Stack{}
	fmt.Printf("Type of variable s: %T\n", s)
	fmt.Println("----Pushing Test Data----")

	for i := 0; i < 5; i++ {
		s.Push(i)
		printStack(s.Display())
	}

	fmt.Println("----Pop Test Data----")
	for j := 0; j < 5; j++ {
		d := s.Pop()
		if d != nil {
			fmt.Printf("Popped: %v\n", d)
			printStack(s.Display())
		}
	}

	fmt.Printf("Is stack now empty? %v\n", s.IsEmpty())
}
