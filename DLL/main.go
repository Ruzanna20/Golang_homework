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

func (dll *DLL) insertAtBegin(data int) {
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

func (dll *DLL) insertAtEnd(data int) {
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

func (dll *DLL) delete(data int) {
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

func (dll *DLL) countOfNode() int {
	var count int

	current := dll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (dll *DLL) checkNegativeData() bool {
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

func (dll *DLL) sumOfAllData() int {
	var sum int

	current := dll.head
	for current != nil {
		sum += current.data
		current = current.next
	}
	return sum
}

func (dll *DLL) sumAtOddIndexes() int {
	var sum int

	current := dll.head
	for current != nil && current.next != nil {
		sum += current.next.data
		current = current.next.next
	}
	return sum
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

	dll.insertAtBegin(10)
	dll.insertAtBegin(20)
	dll.insertAtBegin(30)

	dll.insertAtEnd(50)
	dll.insertAtEnd(60)
	dll.insertAtEnd(70)

	dll.print()

	dll.delete(10)
	dll.print()
	dll.delete(70)
	dll.print()

	count := dll.countOfNode()
	fmt.Println("Count of node in the DLL:", count)

	negativedata := dll.checkNegativeData()
	fmt.Println("DLL has the negative data:", negativedata)

	sumofalldata := dll.sumOfAllData()
	fmt.Println("Sum of the all data in the DLL:", sumofalldata)

	sumofoddindexes := dll.sumAtOddIndexes()
	fmt.Println("Sum of odd index nodes in DLL:", sumofoddindexes)
}
