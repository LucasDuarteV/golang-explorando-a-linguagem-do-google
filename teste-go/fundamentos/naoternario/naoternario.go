package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "reprovado"
}

func main() {
	fmt.Println(obterResultado(10))
}
