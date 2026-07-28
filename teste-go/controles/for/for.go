package main

import "fmt"

func main(){
	i := 1
	for i <= 10{
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++{
		fmt.Printf(" %d " , i)
	}
	fmt.Println("-")
	fmt.Println("Misturando...")

	for i := 1; i < 10; i++{
		if i % 2 == 0 {
			fmt.Println("É impar!")
		} else{
			fmt.Println("É par!")
		}
	}
}
