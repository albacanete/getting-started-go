package main

import (
	"fmt"
	"log"
)

func main() {
	var numF float32

	fmt.Println("Enter a float number: ")

	_, err := fmt.Scan(&numF)
	if err != nil {
		log.Fatal(err)
	}

	var numI int32
	numI = int32(numF)
	fmt.Printf("Its truncated value is %d \n", numI)
}
