package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func importFile(file string) ([][]int, error) {
	var err error
	var mySlice = make([][]int, 0)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		fields := strings.Fields(line)
		temp := make([]int, 0)
		for _, v := range fields {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			temp = append(temp, n)
		}
		if len(temp) != 0 {
			mySlice = append(mySlice, temp)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(temp) != len(mySlice[0]) {
			return nil, errors.New("Wrong number of elements!")
		}
	}
	return mySlice, nil
}

func validPuzzle(sl [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			iEl := i * 3
			jEl := j * 3
			mySlice := []int{0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					bigI := iEl + k
					bigJ := jEl + m
					val := sl[bigI][bigJ]
					if val > 0 && val < 10 {
						if mySlice[val-1] == 1 {
							fmt.Println("Appeared 2 times:", val)
							return false
						} else {
							mySlice[val-1] = 1
						}
					} else {
						fmt.Println("Invalid value")
						return false
					}
				}
			}
		}
	}
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum += sl[i][j]
		}
		if sum != 45 {
			return false
		}
	}
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum += sl[j][i]
		}
		if sum != 45 {
			return false
		}
	}
	return true
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Give me one more arguments")
		return
	}
	file := arguments[1]
	mySlice, err := importFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if validPuzzle(mySlice) {
		fmt.Println("correct sudoku puzzle")
	} else {
		fmt.Println("incorrect sudoku puzzle")
	}
}
