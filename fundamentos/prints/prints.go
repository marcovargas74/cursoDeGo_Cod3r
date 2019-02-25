package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Mesma")
	fmt.Println(" linha.")
	//--------------------------------------
	x := 3.141516

	//fmt.Println("O Valor de x é:" + x)
	xs := fmt.Sprint(x)
	fmt.Println("O Valor de x é:" + xs)
	fmt.Println("O Valor de x é:", x)

	fmt.Printf("O Valor de x é %.2f", x)

	//----------------------------------------
	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
