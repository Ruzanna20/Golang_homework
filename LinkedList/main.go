package main

import "fmt"

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

func (sll *SLL) insertAtBegin(data int) {
	newnode := NewNode(data)
	newnode.next = sll.head
	sll.head = newnode
}

func (sll *SLL) insertAtEnd(data int) {
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

func (sll *SLL) delete(data int) {
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

func (sll *SLL) countOfNode() int {
	var count int

	current := sll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (sll *SLL) checkNegativeData() bool {
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

func (sll *SLL) sumOfAllData() int {
	var sum int

	current := sll.head
	for current != nil {
		sum += current.data
		current = current.next
	}
	return sum
}

func (sll *SLL) sumAtOddIndexes() int {
	var sum int

	current := sll.head
	for current != nil && current.next != nil {
		sum += current.next.data
		current = current.next.next
	}
	return sum
}

func (sll *SLL) print() {
	current := sll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := SLL{}
	ll.insertAtBegin(10)
	ll.insertAtBegin(15)
	ll.insertAtBegin(20)

	ll.insertAtEnd(16)
	ll.insertAtEnd(15)
	ll.insertAtEnd(10)
	ll.print()

	ll.delete(55)
	ll.print()

	count := ll.countOfNode()
	fmt.Println("Count of node in the SLL:", count)

	negativedata := ll.checkNegativeData()
	fmt.Println("SLL has the negative data:", negativedata)

	sumofalldata := ll.sumOfAllData()
	fmt.Println("Sum of the all data in the SLL:", sumofalldata)

	sumofoddindexes := ll.sumAtOddIndexes()
	fmt.Println("Sum of odd index nodes in SLL:", sumofoddindexes)
}
