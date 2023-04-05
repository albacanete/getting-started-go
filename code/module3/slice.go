package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	var result = make([]int, 3) // slice of size 3 initialized to 0
	var input string

	var counter int // counter to replace first 3 zeros with values
	counter = 0

	for {
		fmt.Print("Please, introduce a number (or X to finish): ")
		// read input
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input!")
			log.Fatal(err)
		}

		// if input is "X", stop execution
		if input == "X" {
			break
		}

		// if execution continues, convert input to int
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input is not an integer!")
			log.Fatal(err)
		}

		// if we are in the first 3 elements, overwrite value
		// if not, append new value to existing
		if counter < 3 {
			result[counter] = num
		} else {
			result = append(result, num)
		}

		// if we sort the original result, elements may be overwritten by new inputs
		// we create a copy and sort it
		sorted := make([]int, len(result))
		copy(sorted, result)
		sort.Ints(sorted)

		// print sorted slice
		fmt.Println("Result: ", sorted)

		// update counter
		counter++
	}
}
