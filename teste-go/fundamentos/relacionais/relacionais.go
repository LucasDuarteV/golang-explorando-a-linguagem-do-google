package main

import "fmt"

func main() {
	fmt.Println("Igual =")
	fmt.Println(">" , 10 > 2)
	fmt.Println("Menor < ", 2 < 10)

	type Pessoa struct{
		Nome string
	}

	p1:= Pessoa{"Lucas"}
	p2:=Pessoa{"Lucas"}
	fmt.Println("Pessoas: " , p1 == p2)
}
