package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(time.Second)
	fmt.Println("reading ", <-c)
	time.Sleep(time.Second)
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is close!")
	}
}
