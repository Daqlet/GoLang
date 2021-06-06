package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Daqlet", 170, 60})
	mySlice = append(mySlice, aStructure{"Sonya", 170, 50})
	mySlice = append(mySlice, aStructure{"Marat", 10, 1})
	mySlice = append(mySlice, aStructure{"Sharik", 20, 8})
	fmt.Println("0:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
