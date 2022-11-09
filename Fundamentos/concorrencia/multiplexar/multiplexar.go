package main

import (
	"fmt"

	html "github.com/isadoraclarkson12/generator"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.canaldocampo.com.br/", "http://www.asmeduberaba.com.br/home"),
		html.Titulo("https://picpay.com/", "https://www.ifood.com.br/"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
