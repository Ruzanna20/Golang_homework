package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SLL struct {
	head *Node
}

func NewNode(data int, next *Node) *Node {
	return &Node{
		data: data,
		next: next,
	}
}

func (sll *SLL) insertAtBeginInLL(data int) {
	newnode := NewNode(data, nil)
	newnode.next = sll.head
	sll.head = newnode
}

func (sll *SLL) insertAtEndInLL(data int) {
	newnode := NewNode(data, nil)

	if sll.head == nil {
		sll.head = newnode
	}
	current := sll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newnode
}

func (sll *SLL) deleteFromLL(data int) {
	if sll.head == nil {
		fmt.Println("Datark LL:")
	}

	if sll.head.data == data {
		sll.head = sll.head.next
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
	ll.insertAtBeginInLL(10)
	ll.insertAtBeginInLL(15)
	ll.insertAtBeginInLL(20)

	ll.insertAtEndInLL(16)
	ll.insertAtEndInLL(15)
	ll.insertAtEndInLL(10)
	ll.print()

	ll.deleteFromLL(55)
	ll.print()
}
