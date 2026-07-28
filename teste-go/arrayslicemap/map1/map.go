package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[41581541515] = "Maria"
	aprovados[456415645415] = "Pedro"
	aprovados[46415454544] = "Ana"
	fmt.Println(aprovados)

	for cpf , nome := range aprovados{
		fmt.Printf("%s (CPF:%d)\n" , nome , cpf)
	}

}