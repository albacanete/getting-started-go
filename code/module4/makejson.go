package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var name string
	var addr string

	// read name
	fmt.Println("Please, your name:")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error reading name!")
		log.Fatal(err)
	}

	// read address
	fmt.Println("Please, your address:")
	_, err = fmt.Scan(&addr)
	if err != nil {
		fmt.Println("Error reading address!")
		log.Fatal(err)
	}

	// create map
	pMap := map[string]string{
		"name":    name,
		"address": addr,
	}

	// create JSON with Marhsal
	barr, err := json.Marshal(pMap)
	if err != nil {
		fmt.Println("Error creating JSON!")
		log.Fatal(err)
	}

	// print JSON in string format (if not, bytes printed)
	fmt.Println("JSON: ", string(barr))
}
