package main

import "fmt"

func namesMinMax(x, y int) (min, max int) {
	if x > y {
		max, min = x, y
	} else {
		min, max = x, y
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		max, min = x, y
	} else {
		min, max = x, y
	}
	return min, max
}

func main() {
	var a1, a2 int
	fmt.Scanln(&a1, &a2)
	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)
	fmt.Println(namesMinMax(a1, a2))
	min, max = namesMinMax(a1, a2)
	fmt.Println(min, max)
}
