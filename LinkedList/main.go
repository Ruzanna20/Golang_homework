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

func (sll *SLL) insertAtBeginInSLL(data int) {
	newnode := NewNode(data)
	newnode.next = sll.head
	sll.head = newnode
}

func (sll *SLL) insertAtEndInSLL(data int) {
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

func (sll *SLL) deleteFromSLL(data int) {
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

func (sll *SLL) countOfNodeInSLL() int {
	var count int

	current := sll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (sll *SLL) checkNegativeDataInSLL() bool {
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

func (sll *SLL) sumOfAllDataInSLL() int {
	var sum int

	current := sll.head
	for current != nil {
		sum += current.data
		current = current.next
	}
	return sum
}

func (sll *SLL) sumAtOddIndexesInSLL() int {
	var sum int
	var arrofodd []int

	current := sll.head
	for current != nil {
		arrofodd = append(arrofodd, current.data)
		current = current.next
	}

	for i, val := range arrofodd {
		if i%2 == 1 {
			sum += val
		}
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
	ll.insertAtBeginInSLL(10)
	ll.insertAtBeginInSLL(15)
	ll.insertAtBeginInSLL(20)

	ll.insertAtEndInSLL(16)
	ll.insertAtEndInSLL(15)
	ll.insertAtEndInSLL(10)
	ll.print()

	ll.deleteFromSLL(55)
	ll.print()

	count := ll.countOfNodeInSLL()
	fmt.Println("Count of node in the SLL:", count)

	negativedata := ll.checkNegativeDataInSLL()
	fmt.Println("SLL has the negative data:", negativedata)

	sumofalldata := ll.sumOfAllDataInSLL()
	fmt.Println("Sum of the all data in the SLL:", sumofalldata)

	sumofoddindexes := ll.sumAtOddIndexesInSLL()
	fmt.Println("Sum of odd index nodes in SLL:", sumofoddindexes)
}
