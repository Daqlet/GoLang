package main

import (
	"context"
	"fmt"
	"time"
)

var (
	c1 = context.Background()
	c2 = context.Background()
	c3 = context.Background()
)

func sleep(c *context.CancelFunc) {
	time.Sleep(4 * time.Second)
	(*c)()
}

func f1(t int) {
	c1, cancel := context.WithCancel(c1)
	defer cancel()
	go sleep(&cancel)

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

func f2(t int) {
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()
	go sleep(&cancel)
	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func f3(t int) {
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go sleep(&cancel)
	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	t := 9
	f1(t)
	f2(t)
	f3(t)
}
