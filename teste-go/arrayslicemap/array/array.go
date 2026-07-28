package main

import "fmt"

func main() {
	var notas [3]float64

	notas[1], notas[0], notas[2] = 7, 9, 6

	fmt.Println(notas[0], notas[1], notas[2])

	total := 0.0

	for i := 0; i < len(notas); i++{
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Println("________________")
	fmt.Printf("Média %.2f\n" , media)
}
