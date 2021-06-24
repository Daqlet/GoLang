package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	sum := 0
	sumChan := make(chan int, n)
	w := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		w.Add(1)
		go func(i int) {
			sumChan <- i
			w.Done()
		}(i)
	}
	for i := 0; i < n; i++ {
		w.Add(1)
		go func() {
			x := <-sumChan
			sum += x
			fmt.Println(x)
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(sum)
}
