package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
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

func lookup(t *Node, v int) bool {
	for t != nil {
		if t.Value == v {
			return true
		}
		t = t.Next
	}
	return false
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 4)
	addNode(root, 5)
	addNode(root, 4)
	addNode(root, 5)
	addNode(root, 4)
	addNode(root, 55)
	addNode(root, 48)
	addNode(root, 5)
	traverse(root)
	if lookup(root, 55) {
		fmt.Println("Node exists")
	} else {
		fmt.Println("Node doesn`t exists")
	}
	if lookup(root, 56) {
		fmt.Println("Node exists")
	} else {
		fmt.Println("Node doesn`t exists")
	}
}
