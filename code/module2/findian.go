package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var input string

	fmt.Println("Enter a string: ")

	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	input = strings.ToLower(input)   // all to lower cases
	input = strings.TrimSpace(input) // remove spaces

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
