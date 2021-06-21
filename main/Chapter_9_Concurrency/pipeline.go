package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

var CloseA = false
var Data = make(map[int]bool)
var signal chan struct{}

func first(min, max int, in chan<- int) {
	for {
		if CloseA {
			close(in)
			return
		}
		select {
		case <-signal:
		case in <- random(min, max):
		}
		//in <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for i := range in {
		fmt.Printf("%d ", i)
		_, ok := Data[i]
		if ok {
			CloseA = true
			signal <- struct{}{}
			break
		}
		Data[i] = true
		out <- i
	}
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	fmt.Println("Sum:", sum)
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Need more arguments")
		return
	}
	min, _ := strconv.Atoi(args[1])
	max, _ := strconv.Atoi(args[2])
	A := make(chan int)
	B := make(chan int)
	signal = make(chan struct{})
	rand.Seed(time.Now().Unix())
	go first(min, max, A)
	go second(B, A)
	third(B)
}
