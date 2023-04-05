package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// create struct of type Person
	type Person struct {
		firstName string
		lastName  string
	}

	// create empty slice of type Person
	var people = make([]Person, 0)

	// open file with names of people
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading address!")
		log.Fatal(err)
	}

	// create scanner that will read through the file
	scanner := bufio.NewScanner(f)
	// read maximum 20 characters per line
	barr := make([]byte, 20)
	scanner.Buffer(barr, 20)
	// read every line
	for scanner.Scan() {
		// read and split first name and last name
		pers := scanner.Text()
		fullName := strings.Split(pers, " ")

		// create new person and assign read values
		var p Person
		p.firstName = fullName[0]
		p.lastName = fullName[1]

		// append new person to slice of all people
		people = append(people, p)
	}
	// close file
	f.Close()

	// iterate through slice and print people's name
	for i := 0; i < len(people); i++ {
		fmt.Printf("Person #%d \n", i)
		fmt.Printf("Name: %s \n", people[i].firstName)
		fmt.Printf("Last name: %s \n \n", people[i].lastName)
	}
}
