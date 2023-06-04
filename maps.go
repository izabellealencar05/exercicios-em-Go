package main

import "fmt"

func main() {
	idades := make(map[string]uint8)
	idades["tiago"] = 31
	idades["iza"] = 18
	idades["vh"] = 26

	fmt.Println(idades)
}
