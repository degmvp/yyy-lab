package main

import (
	"fmt"
	"strings"
)

type regPessoa struct {
	nome      string
	altura    float64
	pesoideal float64
}

func main() {
	pessoas := make([]regPessoa, 0, 10) // Slice de pessoa
	var nm string
	var al float64

	for {
		fmt.Println("Digite o nome da pessoa (ou 'FIM' para encerrar):")
		fmt.Scanln(&nm)
		if strings.ToUpper(nm) == "FIM" {
			break
		}

		fmt.Println("Digite a altura da pessoa (em metros):")
		_, err := fmt.Scanln(&al)
		if err != nil {
			fmt.Println("Erro ao ler a altura. Certifique-se de digitar um número válido.")
			continue
		}

		pi := 72.7*al - 58.0 // Cálculo do peso ideal
		pessoa := regPessoa{nome: nm, altura: al, pesoideal: pi}
		pessoas = append(pessoas, pessoa)
	}

	fmt.Println("\nLista de pessoas:")
	for i, v := range pessoas {
		fmt.Printf("%d: Nome: %s, Altura: %.2f m, Peso Ideal: %.2f kg\n", i, v.nome, v.altura, v.pesoideal)
	}
}