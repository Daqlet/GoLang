package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

/*
type Record struct {
	Name string
	Surname string
	Tel []Telephone
}
type Telephone struct {
	Mobile bool
	Number string
}
*/
func loadFromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Give me one more arguments!")
		return
	}
	filename := arguments[1]
	record := Record{}
	err := loadFromXML(filename, &record)
	if err != nil {
		panic(err)
	}
	fmt.Println(record)
}
