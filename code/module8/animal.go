package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// data structure of type animal
type AnimalData struct {
	food       string
	locomotion string
	sound      string
}

// create Animal interface with actions
type Animal interface {
	Eat()
	Move()
	Speak()
}

// return animal's eaten food
func (a AnimalData) Eat() {
	fmt.Println(a.food)
}

// return animal's locomotion method
func (a AnimalData) Move() {
	fmt.Println(a.locomotion)
}

// return animal's spoken sound
func (a AnimalData) Speak() {
	fmt.Println(a.sound)
}

func main() {
	animals := map[string]AnimalData{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	var anInt Animal

	// create reader
	br := bufio.NewReader(os.Stdin)
	for {
		// print prompt
		fmt.Print("> ")
		input, _, err := br.ReadLine()
		if err != nil {
			fmt.Println("Error reading input!")
			log.Fatal(err)
		}

		// separate input words and check command
		in := strings.Split(string(input), " ")
		commmand := in[0]
		switch commmand {
		case "newanimal":
			name := in[1]
			anType := in[2]

			// copy attributes from existing animal
			animals[name] = animals[anType]
			fmt.Println("Created it!")
		case "query":
			name := in[1]
			action := in[2]

			// use interface to get specific animal actions
			anInt = animals[name]

			switch action {
			case "eat":
				anInt.Eat()
			case "move":
				anInt.Move()
			case "speak":
				anInt.Speak()
			default:
				println("Wrong action")
			}
		default:
			fmt.Println("Wrong command")
		}
	}
}
