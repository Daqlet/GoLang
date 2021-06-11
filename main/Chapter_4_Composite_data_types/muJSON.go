package main

import (
	"encoding/json"
	"fmt"
)

//dublicate in the file readJSON.go
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
func main() {
	myRecord := Record{
		Name:    "Daqlet",
		Surname: "Asuov",
		Tel: []Telephone{Telephone{Mobile: true, Number: "1234-567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc-567"},
		},
	}
	rec, err := json.Marshal(&myRecord)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rec))
	var unRec Record
	err = json.Unmarshal(rec, &unRec)
	if err != nil {
		panic(err)
	}
	fmt.Println(unRec)
}
