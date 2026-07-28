package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota < 7 {
		fmt.Println("Você foi reprovado!")
	} else {
		fmt.Println("Você foi aprovado!")
	}
}

func main() {
	imprimirResultado(5)
}