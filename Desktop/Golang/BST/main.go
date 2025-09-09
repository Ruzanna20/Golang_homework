package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

type BST struct {
	root *Node
}

func (bst *BST) insert(data int) {
	newNode := NewNode(data)
	if bst.root == nil {
		bst.root = newNode
		return
	}

	current := bst.root
	for current != nil {
		if data < current.data {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
		}
	}
}

func (bst *BST) search(data int) (int, string) {
	current := bst.root
	for current != nil {
		if current.data == data {
			return current.data, ": Elementy gtnvec"
		} else if data < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}
	return 0, ": Elementy chka"
}

func (bst *BST) inOrder(node *Node) {
	if node != nil {
		bst.inOrder(node.left)
		fmt.Print(node.data, " ")
		bst.inOrder(node.right)
	}
}

func main() {
	bst := BST{}

	bst.insert(8)
	bst.insert(3)
	bst.insert(1)
	bst.insert(10)
	bst.insert(6)
	bst.insert(14)

	bst.inOrder(bst.root)
	fmt.Println()

	data, comment := bst.search(1)
	fmt.Printf("%v %v", data, comment)
}
