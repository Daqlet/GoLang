package main

import (
	"fmt"
	"time"
)

func print() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func main() {
	go print()
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(time.Second)
}
