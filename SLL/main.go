package main

import (
	"fmt"
	"myproject/Interfaces"
)

type Node struct {
	data int
	next *Node
}

type SLL struct {
	head *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

func (sll *SLL) InsertAtBegin(data int) {
	newnode := NewNode(data)
	newnode.next = sll.head
	sll.head = newnode
}

func (sll *SLL) InsertAtEnd(data int) {
	newnode := NewNode(data)

	if sll.head == nil {
		sll.head = newnode
		return
	}
	current := sll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newnode
}

func (sll *SLL) Delete(data int) {
	if sll.head == nil {
		fmt.Println("Datark LL:")
		return 
	}

	if sll.head.data == data {
		sll.head = sll.head.next
		return 
	}

	current := sll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return 
		}
		current = current.next
	}
}

func (sll *SLL) CountOfNode() int {
	if sll.head == nil {
		return 0
	}

	var count int
	current := sll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (sll *SLL) CheckNegativeData() bool {
	if sll.head == nil {
		return false
	}

	var negativedata bool
	current := sll.head
	for current != nil {
		if current.data < 0 {
			negativedata = true
			break
		}
		current = current.next
	}
	return negativedata
}

func (sll *SLL) SumOfAllData() int {
	if sll.head == nil {
		return 0
	}

	var sum int
	current := sll.head
	for current != nil {
		sum += current.data
		current = current.next
	}
	return sum
}

func (sll *SLL) SumAtOddIndexes() int {
	if sll.head == nil {
		return 0
	}

	var sum int
	current := sll.head
	for current != nil && current.next != nil {
		sum += current.next.data
		current = current.next.next
	}
	return sum
}

func (sll *SLL) Print() {
	current := sll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func createSLL(sll interfaces.LL) {
	// Insert elements at the begin
	sll.InsertAtBegin(10)
	sll.InsertAtBegin(15)
	sll.InsertAtBegin(20)

	// Insert elements at the end
	sll.InsertAtEnd(16)
	sll.InsertAtEnd(15)
	sll.InsertAtEnd(10)

	// Print list
	sll.Print()

	// Delete elements 
	sll.Delete(55)

	// Print list after deletion
	sll.Print()
}

func main() {
	sll := &SLL{}
	createSLL(sll)

	fmt.Println("Node count:", sll.CountOfNode())
	fmt.Println("Has negative data:", sll.CheckNegativeData())
	fmt.Println("Sum of all data:", sll.SumOfAllData())
	fmt.Println("Sum at odd indexes:", sll.SumAtOddIndexes())


}
