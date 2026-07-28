package main

import (
	"fmt"
	"time"
)

func falePessoa(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteraçãp %d)\n",pessoa,texto,i+1)
	}
}

func main(){
	falePessoa("Maria","Pq flw comigo?",3)
	falePessoa("João","Só posso falar depois de vc!",3)
}