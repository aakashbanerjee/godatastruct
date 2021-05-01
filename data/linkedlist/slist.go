package linkedlist

import (
	"fmt"
)

//Node represents a single node of a singly linked list.
//A single node in the linked list can hold an item of any type.
type Node struct {
	Item interface{}
	next *Node
}

//LinkedList represents a singly linked lists with a head Node and length of the list
type LinkedList struct {
	head *Node
	len  int
}

//Display prints the items in the liist
func (l *LinkedList) Display() {
	if l.head == nil {
		fmt.Println("Linked List is empty")
	} else {
		currnode := l.head
		for currnode != nil {
			fmt.Printf("Node Item: %v Node Next: %v\n", currnode.Item, currnode.next)
			currnode = currnode.next
		}
	}
}

//InsertHead inserts a Node and makes it the Head of the linked List
func (l *LinkedList) InsertHead(n *Node) {
	headNode := l.head
	if headNode == nil {
		l.head = n
	} else {
		n.next = headNode
		l.head = n
	}
	l.len++
}

//InsertTail inserts a Node at the end of the existing Linked List
func (l *LinkedList) InsertTail(n *Node) {
	headNode := l.head
	if headNode == nil {
		l.head = n
	} else {
		currnode := headNode
		for currnode.next != nil {
			currnode = currnode.next
		}
		currnode.next = n
	}
	l.len++
}
