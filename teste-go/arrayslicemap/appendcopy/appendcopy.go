package main

import "fmt"

func main() {
	var slice []int
	s1 := append(slice, 1 ,2,3)
	fmt.Println(s1)

	s2 := make([]int,2)
	copy(s2,s1)

	fmt.Println(s2)
}	
