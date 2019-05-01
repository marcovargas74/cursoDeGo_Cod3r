package main

import "fmt"

func main() {
	funcsSalario := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	funcsSalario["Rafael Filho"] = 1350.0
	delete(funcsSalario, "Inexistente")

	for nome, salario := range funcsSalario {
		fmt.Println(nome, salario)
	}
}
