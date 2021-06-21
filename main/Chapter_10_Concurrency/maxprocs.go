package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}
