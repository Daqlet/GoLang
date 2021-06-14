package main

import "fmt"

func returnFunction() func() int {
	i := 0
	return func() int {
		i++
		return i*i
	}
}

func main() {
	i := returnFunction()
	j := returnFunction()

	fmt.Println("i1:", i())
	fmt.Println("i2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j1:", j())
	fmt.Println("i3:", i())
}
