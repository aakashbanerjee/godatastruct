package main

import (
	"fmt"

	"github.com/godatastruct/stack"
)

func main() {
	var s stack.Stack
	s.Push(1)
	printStack(s.Display())
	s.Push(2)
	printStack(s.Display())
	s.Push(3)
	printStack(s.Display())
	s.Pop()
	printStack(s.Display())
	fmt.Println("Is Stack Empty: ", s.IsEmpty())
}

func printStack(st []int, t int) {
	fmt.Printf("Stack Elements: %v\t \n", st)
	fmt.Printf("Stack Top: %v\n", t-1)
}
