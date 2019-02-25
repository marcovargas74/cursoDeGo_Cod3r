package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int
	var f *string
	var g *byte

	fmt.Printf("Zeros: %v %v %v %q %v %v %v ", a, b, c, d, e, f, g)
}
