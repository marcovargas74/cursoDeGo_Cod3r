package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"Jose Joao": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	//delete(funcsPorLetra, "P")
	delete(funcsPorLetra["G"], "Guga Pereira")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
