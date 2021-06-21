package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Numbers of goroutines")
	flag.Parse()
	fmt.Printf("Going to create %d goroutines\n", *n)
	count := *n
	var waitGroup sync.WaitGroup
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("Exiting...")
}
