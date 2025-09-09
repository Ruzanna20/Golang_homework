package main

import "fmt"

type Node struct {
	data   int
	left   *Node
	right  *Node
	height int
}

func NewNode(data int) *Node {
	return &Node{
		data:   data,
		height: 1,
	}
}

type AVL struct {
	root *Node
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func balance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func rightRotate(root *Node) *Node {
	l := root.left
	lr := l.right

	

	l.right = root
	root.left = lr

	root.height = 1 + max(height(root.left), height(root.right))
	l.height = 1 + max(height(l.left), height(l.right))

	return l
}

func leftRotate(root *Node) *Node {
	r := root.right
	rl := r.left

	r.left = root
	root.right = rl

	root.height = 1 + max(height(root.left), height(root.right))
	r.height = 1 + max(height(r.left), height(r.right))

	return r
}

func (avl *AVL) insert(data int, node *Node) *Node {
	newNode := NewNode(data)

	if node == nil {
		node = newNode
		return node
	}

	if data < node.data {
		node.left = avl.insert(data,node.left)
	} else if data > node.data {
		node.right = avl.insert(data,node.right)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left),height(node.right))
	balance := balance(node)

	if balance > 1 && data < node.left.data {
		return rightRotate(node)
	}
	if balance > 1 && data > node.left.data {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}
	if balance < -1 && data > node.right.data {
		return leftRotate(node)
	}
	if balance < -1 && data < node.right.data {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}
	return node
}

func (avl *AVL) inOrder(node *Node) {
	if node != nil {
		avl.inOrder(node.left)
		fmt.Print(node.data, " ")
		avl.inOrder(node.right)
	}
}

func main() {
	avl := AVL{}

	avl.root = avl.insert(11,avl.root)
	avl.root = avl.insert(12,avl.root)
	avl.root = avl.insert(10,avl.root)
	avl.root = avl.insert(8,avl.root)
	avl.root = avl.insert(9,avl.root)
	
	avl.inOrder(avl.root)
}