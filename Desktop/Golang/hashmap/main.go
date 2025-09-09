package main

import "fmt"

type Node struct {
	key   string
	value int
	next  *Node
}

type Hashmap struct {
	size    int
	buckets []*Node
}

func NewHashMap(size int) *Hashmap {
	buckets := make([]*Node, size)
	return &Hashmap{
		size:    size,
		buckets: buckets,
	}
}

func NewNode(key string, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func HashFunc(key string, size int) int {
	hash := len(key) % size
	return hash
}

func (hm *Hashmap) SetVal(key string, value int) {
	bucketIndex := HashFunc(key, hm.size)
	newNode := NewNode(key, value)

	if hm.buckets[bucketIndex] == nil {
		hm.buckets[bucketIndex] = newNode
		return
	}

	current := hm.buckets[bucketIndex]
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (hm *Hashmap) GetVal(key string) int {
	bucketIndex := HashFunc(key, hm.size)

	if hm.buckets[bucketIndex] == nil {
		return -1
	}

	current := hm.buckets[bucketIndex]
	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return -1
}

func (hm *Hashmap) Print() {
	fmt.Println("\nHashmap:")
	for index, node := range hm.buckets {
		fmt.Printf("Bucket %v: ", index)
		current := node
		for current != nil {
			fmt.Printf("(%v : %v )-> ", current.key, current.value)
			current = current.next
		}
		fmt.Println("nil")
	}
}

func main() {
	hm := NewHashMap(5)

	hm.SetVal("Anna", 1)
	hm.SetVal("David", 2)
	hm.SetVal("Anna", 3)
	hm.SetVal("David", 2)
	hm.SetVal("Ani", 5)

	fmt.Println("Getting Value:", hm.GetVal("David"))
	fmt.Println("Getting Value:", hm.GetVal("Ruzanna"))

	hm.Print()
}
