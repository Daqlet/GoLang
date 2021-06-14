package main

import "fmt"

func main() {
	var myInt interface{} = 123

	i, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", i)
	}
	j, ok := myInt.(float64)
	if ok {
		fmt.Println(j)
	} else {
		fmt.Println("Failing without panicing")
	}
	i = myInt.(int)
	fmt.Println(i)

	k := myInt.(bool)
	fmt.Println(k)
}
