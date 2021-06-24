package main

import (
	"fmt"
	"sync"
)

var (
	n   int = 10
	sum int = 0
	m   sync.Mutex
	ch  = make(chan int, 10)
)

func fillChan(w *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			ch <- i
		}(i)
	}
}

func findSum(w *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			m.Lock()
			sum += <-ch
			m.Unlock()
		}()
	}
}

func main() {
	w := sync.WaitGroup{}
	fillChan(&w)
	findSum(&w)
	w.Wait()
	fmt.Println(sum)
}
