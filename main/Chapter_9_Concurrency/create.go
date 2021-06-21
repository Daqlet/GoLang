package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()
	fmt.Printf("Going to create %d goroutines\n", *n)
	count := *n
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}
