package main

import (
	"fmt"

	"github.com/godatastruct/data/stack"
)

func main() {
	s := &stack.Stack{}
	fmt.Printf("Type of variable s: %T\n", s)
	s.Push("aakash")
	printStack(s.Display())
	s.Push(2)
	printStack(s.Display())
	s.Push(3.14)
	printStack(s.Display())
	d := s.Pop()
	if d != nil {
		fmt.Printf("Popped: %v\n", d)
	}

	printStack(s.Display())
	fmt.Println("Is Stack Empty: ", s.IsEmpty())
	d = s.Pop()
	if d != nil {
		fmt.Printf("Popped: %v\n", d)
	}
	printStack(s.Display())
	fmt.Println("Is Stack Empty: ", s.IsEmpty())
	d = s.Pop()
	if d != nil {
		fmt.Printf("Popped: %v\n", d)
	}
	printStack(s.Display())
	fmt.Println("Is Stack Empty: ", s.IsEmpty())
	d = s.Pop()
	if d != nil {
		fmt.Printf("Popped: %v\n", d)
	}
	printStack(s.Display())
	fmt.Println("Is Stack Empty: ", s.IsEmpty())
}

func printStack(st []interface{}, t int) {
	fmt.Printf("Stack Elements: %v\t \n", st)
	//fmt.Printf("Stack Top: %v\n", t)
}
