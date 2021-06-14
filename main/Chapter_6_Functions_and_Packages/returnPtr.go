package main

import "fmt"

func returnPtr(x int) *int {
	y := x*x
	return &y
}

func main() {
	sq := returnPtr(10)
	fmt.Println("Ans:", *sq)
	fmt.Println("Adress:", sq)
}
