package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro
// para um struct Livro com campos título e autor, e
// altere o título do livro para "Desconhecido" se o
// autor for "Anônimo".
type Livro struct {
	Titulo string
	Autor  string
}

func ptr(livro *Livro) {
	if livro.Autor == "anônimo" {
		livro.Titulo = "desconhecido"
	}
}
func main() {
	meuLivro := Livro{Titulo: "livro A", Autor: "autor anônimo"}
	ptr(&meuLivro)
	fmt.Println(meuLivro)
}
