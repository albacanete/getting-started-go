package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// data structure of type animal
type Animal struct {
	food       string
	locomotion string
	sound      string
}

// return animal's eaten food
func Eat(a Animal) {
	fmt.Println(a.food)
}

// return animal's locomotion method
func Move(a Animal) {
	fmt.Println(a.locomotion)
}

// return animal's spoken sound
func Speak(a Animal) {
	fmt.Println(a.sound)
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

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

		// separate both input words
		in := strings.Split(string(input), " ")
		name := in[0]
		action := in[1]

		// print bt type of animal
		switch action {
		case "eat":
			Eat(animals[name])
		case "move":
			Move(animals[name])
		case "speak":
			Speak(animals[name])
		default:
			println("Wrong action")
		}
	}

}
