package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma :",a + b)
	fmt.Println("Subração: " , a - b)
	fmt.Println("Divisão: ", a/b)
	fmt.Println("Módulo: ", a%b)

	fmt.Println("Maior que: " , math.Max(float64(a) , float64(b)))
	fmt.Println("Menor que: " , math.Min(float64(a) , float64(b)))
}
