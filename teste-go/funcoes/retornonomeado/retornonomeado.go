package main

import "fmt"

func trocar(p1, p2 int) (primeiro int, segundo int) {
	segundo = p1
	primeiro = p2
	return
}

func main() {
	fmt.Println(trocar(2,3))
}
