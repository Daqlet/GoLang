package main

import (
	"fmt"
	"os"
)

func varFunc(input ...string) {
	fmt.Println(input)
}

func oneByOne(s string, n ...int) int {
	fmt.Println(s)
	sum := 0
	for i, v := range n {
		fmt.Println(i, v)
		sum += i
	}
	n[0] = -1000
	return sum
}

func main() {
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByOne("Adding numbers...", 1, 5, -8, 5, 9 ,-4 ,-6 ,6)
	fmt.Println("sum:", sum)
	s := []int{1, 2, 3}
	fmt.Println(oneByOne("Adding numers...", s...))
	fmt.Println(s)
}
