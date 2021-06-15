package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int
type t2 int

type a struct {
	X int
	Y float64
	S string
}

func (a1 a) compareStruct(a2 a) bool {
	r1 := reflect.ValueOf(&a1).Elem()
	r2 := reflect.ValueOf(&a2).Elem()
	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Println("Type to examine:", t)
	for i := 0; i < r.NumMethod(); i++ {
		m := r.Method(i).Type()
		fmt.Println(t.Method(i).Name, "--->", m)
	}
}

func main() {
	c := t1(100)
	b := t2(200)
	fmt.Println("The type of c is", reflect.TypeOf(&c))
	fmt.Println("The type of b is", reflect.TypeOf(&b))
	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Println("The type of r is", reflect.TypeOf(r))
	a1 := a{1, 2.1, "A1"}
	a2 := a{1, -2, "A2"}
	if a1.compareStruct(a1) {
		fmt.Println("Equal!")
	}
	if !a1.compareStruct(a2) {
		fmt.Println("Not Equal!")
	}
	var f *os.File
	printMethods(f)
}
