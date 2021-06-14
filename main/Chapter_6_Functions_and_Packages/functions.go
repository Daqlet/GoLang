package main

import (
	"fmt"
	"os"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return 2*x, x*x
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need more arguments")
		return
	}
	y, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	square := func(s int) int {
		return s*s
	}
	fmt.Println("The square of", y, "is", square(y))
	double := func(d int) int {
		return 2*d
	}
	fmt.Println("The double of", y, "is", double(y))
	fmt.Println(doubleSquare(y))
	d, s := doubleSquare(y)
	fmt.Println(d, s)
}
