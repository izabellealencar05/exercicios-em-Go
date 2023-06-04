package main

import "fmt"

func main() {
	//slices, pois n tem um tamanho fixo
	nomes := []string{"izabelle", "victor", "iza", "waldemar"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	//"len" = tamanho da lista, "cap" = capacidade da lista

	nomes = append(nomes, "clara") // "append" adiciona um adicional
	fmt.Println(nomes, len(nomes), cap(nomes))

}
