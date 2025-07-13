package main

import (
	"fmt"
	"myproject/Interfaces"
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

func (dll *DLL) InsertAtBegin(data int) {
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

func (dll *DLL) InsertAtEnd(data int) {
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

func (dll *DLL) Delete(data int) {
	if dll.head == nil {
		fmt.Println("Datark DLL:")
		return 
	}

	if dll.head.data == data {
		if dll.head.next != nil {
			dll.head = dll.head.next
			dll.head.prev = nil
		} else {
			dll.head = nil
			dll.tail = nil
		}
		return 
	}

	current := dll.head
	for current.next != nil {
		if current.next.data == data {
			if current.next.next != nil {
				current.next = current.next.next
				current.next.prev = current
			} else {
				current.next = nil
				dll.tail = current
			}
			return
		}
		current = current.next
	}
}

func (dll *DLL) CountOfNode() int {
	if dll.head == nil {
		return 0
	}

	var count int
	current := dll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (dll *DLL) CheckNegativeData() bool {
	if dll.head == nil {
		return false
	}

	var negativedata bool
	current := dll.head
	for current != nil {
		if current.data < 0 {
			negativedata = true
			break
		}
		current = current.next
	}
	return negativedata
}

func (dll *DLL) SumOfAllData() int {
	if dll.head == nil {
		return 0
	}

	var sum int
	current := dll.head
	for current != nil {
		sum += current.data
		current = current.next
	}
	return sum
}

func (dll *DLL) SumAtOddIndexes() int {
	if dll.head == nil {
		return 0
	}

	var sum int
	current := dll.head
	for current != nil && current.next != nil {
		sum += current.next.data
		current = current.next.next
	}
	return sum
}

func (dll *DLL) Print() {
	current := dll.head
	fmt.Printf("nil")
	for current != nil {
		fmt.Printf(" <-> %v", current.data)
		current = current.next
	}
	fmt.Println(" <-> nil")
}

func createDLL(dll interfaces.LL) {
	// Insert elements at the begin
	dll.InsertAtBegin(10)
	dll.InsertAtBegin(20)
	dll.InsertAtBegin(30)

	// Insert elements at the end
	dll.InsertAtEnd(50)
	dll.InsertAtEnd(60)
	dll.InsertAtEnd(70)

	// Print list
	dll.Print()

	// Delete elements 
	dll.Delete(10)
	dll.Delete(70)

	// Print list after deletion
	dll.Print()
}

func main() {
	dll := &DLL{}
	createDLL(dll)

	fmt.Println("Count node of the DLL:", dll.CountOfNode())
	fmt.Println("DLL has negative data:", dll.CheckNegativeData())
	fmt.Println("Sum of node in DLL:", dll.SumOfAllData())
	fmt.Println("Sum of odd nodes in DLL:", dll.SumAtOddIndexes())
}
