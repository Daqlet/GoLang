package aPackage

import "fmt"

func A() {
	fmt.Println("Function A() called!")
}

func B() {
	fmt.Println("Private constant:", privateConstant)
}

const privateConstant = 56464
