package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

var root = new(Node)

func addNode(t *Node, value int) int {
	if root == nil {
		t = &Node{value, nil, nil}
		root = t
		return 0
	}
	if value == t.Value {
		fmt.Println("Node already exists")
		return -1
	}
	if t.Next == nil {
		temp := t
		t.Next = &Node{value, temp, nil}
		return -2
	}
	return addNode(t.Next, value)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}
	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}
	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list")
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookup(t *Node, v int) bool {
	if root == nil {
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookup(t.Next, v)
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 50)
	addNode(root, 100)
	addNode(root, 150)
	traverse(root)
	addNode(root, 40)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}
