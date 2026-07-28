package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for _,num := range numeros {
		fmt.Println(num)
	}
}