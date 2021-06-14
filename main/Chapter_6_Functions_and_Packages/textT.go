package main

import (
	"fmt"
	"html/template"
	"os"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Need the templates file")
		return
	}
	tFile := arguments[1]
	Data := [][]int{{-1, 5}, {5, 6}, {9, 1}, {2, 2}}
	Entries := make([]Entry, 0)
	for _, i := range Data {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			Entries = append(Entries, temp)
		}
	}
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
