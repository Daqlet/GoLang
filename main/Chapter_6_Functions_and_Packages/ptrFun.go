package main

import "fmt"

func sqr(x *float64) float64 {
	return *x * *x
}

func main() {
	var x float64
	fmt.Scanln(&x)
	fmt.Println("sqrt:", sqr(&x))
}