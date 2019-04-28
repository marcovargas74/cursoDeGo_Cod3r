package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { //não tem while em Go
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println("\nFor em uma Linha")
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, "=Par ")
		} else {
			fmt.Print(i, "=ImPar ")
		}
	}

	for {
		//Laço Infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}
}
