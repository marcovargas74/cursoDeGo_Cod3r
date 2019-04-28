package main

import "fmt"

func notaparaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida"
	}

}

func main() {
	fmt.Println(notaparaConceito(9.8))
	fmt.Println(notaparaConceito(6.9))
	fmt.Println(notaparaConceito(2.1))
}