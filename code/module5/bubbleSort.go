package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(a []int, j int) {
	a[j], a[j+1] = a[j+1], a[j]
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	// read integers separated by a space
	fmt.Println("Please input numbers(separate with space):")
	br := bufio.NewReader(os.Stdin)
	a, _, err := br.ReadLine()
	if err != nil {
		fmt.Println("Error reading numbers!")
		log.Fatal(err)
	}

	// get each number separately
	nums := strings.Split(string(a), " ")
	if len(nums) > 10 {
		log.Fatal("There are more than 10 numbers!")
	}

	// create slice to store arrays
	var results = make([]int, 0)

	// convert inputs to ints and add to slice
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Input is not an integer!")
			log.Fatal(err)
		}
		results = append(results, n)
	}

	// sort and print
	BubbleSort(results)
	fmt.Println("Sorted:")
	fmt.Println(results)
}
