package main

import "fmt"

type Pessoa struct {
	nome   string
	idade  int
	status bool
}

func main() {
	p := Pessoa{
		nome:   "tiago",
		idade:  19,
		status: false,
	}
	fmt.Println(p)
	fmt.Println(p.nome, p.idade)
}
