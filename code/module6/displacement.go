package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// Input: acceleration a, initial velocity v0, and initialdisplacement s0
// Output: displacement as a function of time
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	disp := func(t float64) (s float64) {
		s = .5*a*math.Pow(t, 2) + v0*t + s0
		return
	}
	return disp
}

func ReadInput(name string) float64 {
	var input string

	// read variable
	fmt.Printf("Please, enter the %s: \n", name)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Printf("Error reading %s! \n", name)
		log.Fatal(err)
	}

	// convert to float64
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input is not an integer!")
		log.Fatal(err)
	}
	return num
}

func main() {
	// create variables
	var a, v0, s0, t float64

	// read inputs
	a = ReadInput("acceleration")
	v0 = ReadInput("initial velocity")
	s0 = ReadInput("initial displacement")
	t = ReadInput("time interval")

	// calculate displacement
	disp := GenDisplaceFn(a, v0, s0)

	fmt.Printf("After %v seconds the displacement is %v. \n", t, disp(t))
}
