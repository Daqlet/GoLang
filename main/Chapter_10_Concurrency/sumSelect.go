package main

import (
	"fmt"
	"sync"
)

var writeNumber = make(chan int)
var readNumber = make(chan int)

func set(n int) {
	writeNumber <- n
}

func read() int {
	return <-readNumber
}

var sum = 0

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeNumber:
			value = newValue
			sum += value
			fmt.Println(value)
		case readNumber <- value:
		}
	}
}

func main() {
	n := 10
	w := sync.WaitGroup{}
	go monitor()
	for i := 0; i < n; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			set(i)
		}(i)
	}
	w.Wait()
	fmt.Println(sum, read())
}
