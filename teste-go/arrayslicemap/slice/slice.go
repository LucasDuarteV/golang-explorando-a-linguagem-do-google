package main

import "fmt"

func main() {
	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3]
	fmt.Println(a2,s2)

	a3 := [5]int{0,1,2,3,4}
	s3 := a3[2:4]
	fmt.Println(s3)
}
