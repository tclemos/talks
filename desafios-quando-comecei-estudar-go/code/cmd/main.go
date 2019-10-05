package main

import (
	"biblioteca"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Este é o argumento fornecido: %s", os.Args[1])
	} else {
		fmt.Printf("Nenhum argumento foi fornecido")
	}

	algo := biblioteca.CriarAlgo(true)
	algo.Esconder()
}
