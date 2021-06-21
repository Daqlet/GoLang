package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(time.Second * 4):
			fmt.Println("\nTime after!")
		}
	}
}

func main() {
	n := flag.Int("n", 10, "Numbers of random numbers")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d numbers\n", count)
	createNumber := make(chan int)
	end := make(chan bool)
	go gen(0, 2*count, createNumber, end)
	for i := 0; i < count; i++ {
		fmt.Print(<-createNumber, " ")
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Exiting...")
	end <- true
}
