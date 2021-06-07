package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("No error today in returnError()")
	} else {
		fmt.Println(err)
	}
	err = returnError(10, 10)
	/*if err != nil {
		//log.Println(err)
		panic(err)
		os.Exit(10)
	}*/
	if err == nil {
		fmt.Println("No error today in returnError()")
	} else {
		fmt.Println(err)
	}
	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}
