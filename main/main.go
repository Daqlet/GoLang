package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Why I can`t get repos:(\n"
	v4 := "abc"
	fmt.Print(v1, v2, v3, v4)
	fmt.Println()
	fmt.Println(v1, v2, v3, v4)
	fmt.Printf("%s %d %s %s\n", v1, v2, v3, v4)
}
