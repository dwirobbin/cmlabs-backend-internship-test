package helper

// package main

import (
	// "fmt"
	"strconv"
)

func FizzBuzz(from, to int) string {
	fizz := "fizz"
	buzz := "buzz"

	for ; from <= to; from++ {
		if from%15 == 0 {
			return fizz + buzz
		} else if from%5 == 0 {
			return buzz
		} else if from%3 == 0 {
			return fizz
		} else {
			return strconv.Itoa(from)
		}
	}

	return ""
}

// =====Run Program Without Unit Test=====
// func main() {
// 	var from, to = 1, 100

// 	for ; from <= to; from++ {
// 		fmt.Println(FizzBuzz(from, to))
// 	}
// }
