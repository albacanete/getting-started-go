package main

import (
	"fmt"
	"log"
)

func main() {
	var numF float64

	fmt.Println("Enter a float number: ")

	_, err := fmt.Scan(&numF)
	if err != nil {
		log.Fatal(err)
	}

	var numI int
	numI = int(numF)
	fmt.Printf("Its truncated value is %d \n", numI)
}
