package main

import (
	"fmt"
	interfaces "myproject/Interfaces"
)

type Node struct {
	data int
	next *Node
}

type CSLL struct {
	head *Node
}

func newNode(data int) *Node {
	return &Node{
		data: data,
	}
}

func (csll *CSLL) InsertAtBegin(data int) {
	newNode := newNode(data)

	if csll.head == nil {
		csll.head = newNode
		newNode.next = csll.head
		return
	}

	current := csll.head
	for current.next != csll.head {
		current = current.next
	}
	newNode.next = csll.head
	csll.head = newNode
	current.next = csll.head
}

func (csll *CSLL) InsertAtEnd(data int) {
	newNode := newNode(data)

	if csll.head == nil {
		csll.head = newNode
		newNode.next = csll.head
		return
	}

	current := csll.head
	for current.next != csll.head {
		current = current.next
	}
	current.next = newNode
	newNode.next = csll.head
}

func (csll *CSLL) Delete(data int) {
	if csll.head == nil {
		fmt.Println("Datark CSLL:")
		return 
	}

	current := csll.head
	if csll.head.data == data {
		if current.next == csll.head {
			csll.head = nil
		} else {
			for current.next != csll.head {
				current = current.next
			}
			csll.head = csll.head.next
			current.next = csll.head
		}
		return 
	}

	for current.next != csll.head {
		if current.next.data == data {
			current.next = current.next.next
			return 
		}
		current = current.next
	}
}

func (csll *CSLL) Print() {
	if csll.head == nil {
		fmt.Println("Datark CSLL:")
		return
	}

	current := csll.head
	for current.next != csll.head {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
	fmt.Println(current.data, "-> (head)")
}

func (csll *CSLL) CountOfNode() int {
	if csll.head == nil {
		return 0
	}

	count := 1
	current := csll.head
	for current.next != csll.head {
		count++
		current = current.next
	}
	return count
}

func (csll *CSLL) CheckNegativeData() bool {
	if csll.head == nil {
		return false
	}

	var negativedata bool
	if csll.head.data < 0 {
		negativedata = true
		return negativedata
	}

	current := csll.head.next
	for current != csll.head {
		if current.data < 0 {
			negativedata = true
			break
		}
		current = current.next
	}
	return negativedata
}

func (csll *CSLL) SumOfAllData() int {
	if csll.head == nil {
		return 0
	}

	var sum int
	sum += csll.head.data
	current := csll.head.next
	for current != csll.head {
		sum += current.data
		current = current.next
	}
	return sum
}

func (csll *CSLL) SumAtOddIndexes() int {
	if csll.head == nil {
		return 0
	}

	var sum int
	var index int

	current := csll.head
	for {
		if index % 2 == 1 {
			sum += current.data
		}
		current = current.next
		index++
		if current == csll.head {
			break
		}
	}
	return sum
}

func (csll *CSLL) checkCycle() bool {
	if csll.head == nil {
		fmt.Println("Datark CSLL:")
		return false
	}

	var iscycle bool
	fast := csll.head
	slow := csll.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			iscycle = true
			return iscycle
		}
	}
	return iscycle
}

func createCSLL(csll interfaces.LL) {
	// Insert elements at the begin
	csll.InsertAtBegin(17)
	csll.InsertAtBegin(15)

	// Insert elements at the end
	csll.InsertAtEnd(19)
	csll.InsertAtEnd(6)

	// Print list
	csll.Print()

	// Delete from list
	csll.Delete(27)

	// Print after deletion
	csll.Print()
}

func main() {
	csll := &CSLL{}
	createCSLL(csll)


	fmt.Println("Count node of the CSLL:", csll.CountOfNode())
	fmt.Println("CSLL has negative data:", csll.CheckNegativeData())
	fmt.Println("Sum of node in CSLL:", csll.SumOfAllData())
	fmt.Println("Sum of odd nodes in CSLL:", csll.SumAtOddIndexes())

	iscycle := csll.checkCycle()
	fmt.Println("My list is Cycle:", iscycle)
}
