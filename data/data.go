package main

import (
	"fmt"

	"github.com/godatastruct/data/linkedlist"
	"github.com/godatastruct/data/queue"
	"github.com/godatastruct/data/stack"
)

func main() {
	//callStack() //Executes all func calls for Stack
	//callQueue() //Executes all func calls for Queue
	callLinkedList() //Executes all func calls for LinkedList
}

func printStack(st []interface{}, t int) {
	fmt.Printf("Stack Elements: %v Number of elements: %v\n", st, t)
}

func printQueue(q []interface{}, h int) {
	fmt.Printf("Queue Elements: %v Head: %v\n", q, h)
}

func callLinkedList() {
	l := &linkedlist.LinkedList{}
	n := &linkedlist.Node{Item: "D"}
	l.InsertHead(n)
	//fmt.Println(l.Length())
	n = &linkedlist.Node{Item: "C"}
	l.InsertHead(n)
	//fmt.Println(l.Length())
	n = &linkedlist.Node{Item: "B"}
	l.InsertHead(n)
	//fmt.Println(l.Length())
	n = &linkedlist.Node{Item: "E"}
	l.InsertTail(n)
	//fmt.Println(l.Length())
	n = &linkedlist.Node{Item: "A"}
	l.InsertHead(n)
	//fmt.Println(l.Length())
	l.Display()
	l.RemoveHead()
	l.Display()
	l.RemoveTail()
	l.Display()
	fmt.Println(l.Length())
}

func callQueue() {
	fmt.Println("----------QUEUE----------")
	q := &queue.Queue{}
	fmt.Printf("Type of variable q: %T\n", q)
	fmt.Println("----Enqueue Test Data----")
	q.Enqueue(1)
	printQueue(q.Display())
	q.Enqueue(2)
	printQueue(q.Display())
	q.Enqueue(3)
	printQueue(q.Display())
	e, empty := q.Dequeue()
	if empty == true {
		fmt.Println("Queue is Empty")
	}
	fmt.Println(e)
	printQueue(q.Display())
	e, empty = q.Dequeue()
	if empty == true {
		fmt.Println("Queue is Empty")
	}
	fmt.Println(e)
	printQueue(q.Display())
	e, empty = q.Dequeue()
	if empty == true {
		fmt.Println("Queue is Empty")
	}
	fmt.Println(e)
	printQueue(q.Display())
	e, empty = q.Dequeue()
	if empty == true {
		fmt.Println("Queue is Empty")
	}
	fmt.Println(e)
	printQueue(q.Display())
}

func callStack() {
	fmt.Println("----------STACK----------")
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
