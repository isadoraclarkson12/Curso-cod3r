package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func (c carro) getNome() string {
	return c.nome
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 30
	f.turboLigado = true
	fmt.Printf("A ferrari %s está com o turbo ligado? %t", f.getNome(), f.turboLigado)
}
