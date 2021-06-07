package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("You are using", runtime.Compiler, "on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of GPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	fmt.Println("Version:", m1, ".", m2)
}
