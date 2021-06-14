package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func (t *TwoInts) method(a TwoInts) TwoInts {
	temp := TwoInts{t.a + a.a, t.b + a.b}
	return temp
}

func main() {
	a := TwoInts{5, 6}
	b := TwoInts{9, -7}
	fmt.Println(a.method(b))
}
