package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

type DLL struct {
	head *Node
	tail *Node
}

func (dll *DLL) insertAtBeginInDLL(data int) {
	newnode := NewNode(data)

	if dll.head != nil {
		dll.head.prev = newnode
		newnode.next = dll.head
		dll.head = newnode
	} else {
		dll.head = newnode
		dll.tail = newnode
	}
}

func (dll *DLL) insertAtEndInDLL(data int) {
	newnode := NewNode(data)

	if dll.head == nil && dll.tail == nil {
		dll.head = newnode
		dll.tail = newnode
		return
	}

	current := dll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newnode
	newnode.prev = current
	dll.tail = newnode
}

func (dll *DLL) deleteFromDLL(data int) {
	if dll.head == nil {
		fmt.Println("Datark DLL:")
		return
	}

	if dll.head.data == data && dll.head.next != nil {
		dll.head = dll.head.next
		dll.head.prev = nil
		return
	}

	if dll.head.data == data && dll.head.next == nil {
		dll.head = nil
		dll.tail = nil
		return 
	}

	current := dll.head
	for current.next != nil {
		if current.next.data == data && current.next.next != nil {
			current.next = current.next.next
			current.next.prev = current
		} else if current.next.data == data && current.next.next == nil {
			current.next = nil 
			dll.tail = current
			return
		}
		current = current.next
	}
}

func (dll *DLL) print() {
	current := dll.head
	fmt.Printf("nil")
	for current != nil {
		fmt.Printf(" <-> %v", current.data)
		current = current.next
	}
	fmt.Println(" <-> nil")
}

func main() {
	dll := DLL{}

	dll.insertAtBeginInDLL(10)
	dll.insertAtBeginInDLL(20)
	dll.insertAtBeginInDLL(30)

	dll.insertAtEndInDLL(50)
	dll.insertAtEndInDLL(60)
	dll.insertAtEndInDLL(70)

	dll.print()

	dll.deleteFromDLL(10)
	dll.print()
	dll.deleteFromDLL(70)
	dll.print()
}
