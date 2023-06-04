package main

import "fmt"

//Implemente uma função que receba um ponteiro
//para uma struct "Livro" com campos "Título",
//"Autor" e "Preço", e que adicione um desconto
//de 10% no preço do livro.

type Livro struct {
	titulo string
	autor  string
	preco  float64
}

func descontoLivro(l *Livro) {
	l.preco = l.preco * 0.9
}

func main() {
	livro :=
		Livro{titulo: "anxious people",
			autor: "fredrik backman",
			preco: 80.0}
	descontoLivro(&livro)
	fmt.Printf("Livro: %s\nAutor: %s\nPreço com desconto: %.2f\n", livro.titulo, livro.autor, livro.preco)
}
