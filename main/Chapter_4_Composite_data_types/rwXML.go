package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJson := json.NewDecoder(in)
	err = decodeJson.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Give me more arguments, dummy!")
		return
	}
	filename := arguments[1]
	var myRecord Record
	err := loadFromJson(filename, &myRecord)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", myRecord)
	myRecord.Name = "Daqlet"
	xmlData, _ := xml.MarshalIndent(myRecord, "", "	")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData:", string(xmlData))
	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		panic(err)
	}
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("\nJson:", myRecord)
}
