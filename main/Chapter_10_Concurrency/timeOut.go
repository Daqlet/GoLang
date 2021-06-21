package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(time.Second * 5)
		w.Wait()
	}()
	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	n := flag.Int("n", 10, "milliseconds")
	flag.Parse()
	mil := *n
	var w sync.WaitGroup
	w.Add(1)
	fmt.Println("Time out after", mil)
	if timeout(&w, time.Duration(int32(mil))*time.Millisecond) {
		fmt.Println("time out")
	} else {
		fmt.Println("ok")
	}
	w.Done()
	if timeout(&w, time.Duration(int32(mil))*time.Millisecond) {
		fmt.Println("time out")
	} else {
		fmt.Println("ok")
	}
}
