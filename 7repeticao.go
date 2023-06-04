package main

import "fmt"

// Algoritmo para imprimir os números de 1 a 100, substituindo os múltiplos de 3 pela palavra "Fizz" e os múltiplos de 5 pela palavra "Buzz". Para os números que são múltiplos de ambos, utilize a palavra "FizzBuzz":

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
