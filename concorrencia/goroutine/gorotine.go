package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//go fale("Maria", "Pq vc nao fala comigo?", 3)
	//go fale("Joao", "Só posso falar depois de Vc!", 1)

	//go fale("Maria", "Ei...", 500)
	//go fale("Joao", "Opa...", 500)

	go fale("Maria", "Endendi...", 10)
	fale("Joao", "Parabens!", 5)

}
