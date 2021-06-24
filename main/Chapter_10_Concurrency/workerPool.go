package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var size = 1
var client = make(chan Client, size)
var data = make(chan Data, size)

func makeWP(n int) {
	w := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func worker(w *sync.WaitGroup) {
	for c := range client {
		output := Data{c, c.integer * c.integer}
		data <- output
	}
	w.Done()
}

func createJob(n int) {
	for i := 0; i < n; i++ {
		client <- Client{i, i}
	}
	close(client)
}

func main() {
	var n int
	fmt.Scanln(&n)
	go createJob(n)
	finish := make(chan interface{})
	f, err := os.Create("worker.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	var m sync.Mutex
	go func() {
		for d := range data {
			m.Lock()
			f.WriteString(strconv.Itoa(d.job.id) + " " + strconv.Itoa(d.square) + "\n")
			m.Unlock()
		}
		finish <- true
	}()
	makeWP(n)
	fmt.Println(<-finish)
}
