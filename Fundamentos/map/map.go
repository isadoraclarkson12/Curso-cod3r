package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[36452352323] = "Maria"
	aprovados[12343543543] = "Grosbilda"
	aprovados[47556346343] = "Estrobilobaldo"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 47556346343)
	fmt.Println(aprovados)
}
