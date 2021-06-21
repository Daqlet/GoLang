package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (n *NamesFlag) GetNames() []string {
	return n.Names
}

func (n *NamesFlag) String() string {
	return fmt.Sprint(n.Names)
}

func (n *NamesFlag) Set(v string) error {
	if len(n.Names) > 0 {
		return fmt.Errorf("Cannot use names more that once!")
	}
	names := strings.Split(v, ",")
	for _, item := range names {
		n.Names = append(n.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Daqlet", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")
	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)
	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}
	fmt.Println("Remaining command line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
