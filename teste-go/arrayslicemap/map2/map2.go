package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José jõao":     11325.45,
		"Gabriel silva": 15456.78,
		"Pedro Junior":  1200.0,
	}

	funcsESalarios["Rafael filho"] = 1350.0

	for nome, salario := range funcsESalarios {
		fmt.Println(nome,salario)
	}
}
