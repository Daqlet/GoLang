package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Give me more arguments, dummy!")
		return
	}
	filename := arguments[1]
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var parsedData map[string]interface{}
	json.Unmarshal([]byte(fileData), &parsedData)
	for key, value := range parsedData {
		fmt.Println("key:", key, "value:", value)
	}
}
