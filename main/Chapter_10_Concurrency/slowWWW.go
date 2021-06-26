package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, "Serving: %s \n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d \n", delay)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	PORT := ":8001"
	fmt.Println("Using port: ", PORT)
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
