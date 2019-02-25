package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8, uint16, uint32, uint64
	var b byte
	fmt.Println("O bypte é: ", reflect.TypeOf(b))

	//com sinal... uint8, uint16, uint32, uint64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é: ", i1)
	fmt.Println("O tipo de i1 é: ", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabel Unicode(int32)
	fmt.Println("O rune é: ", reflect.TypeOf(i2))
	fmt.Println(i2)
	fmt.Printf("Tabela ascii= 0x%x\n", i2)

	//numeros Reais... float32, float64
	var x float32 = 49
	fmt.Println("O tipo de x é: ", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é: ", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é: ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Marco"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Marco`
	fmt.Println(s2 + "!")
	fmt.Println("O tamanho da string é", len(s2))

	//Char ??? nao existe o tipo char especifico igual ao C
	//var x char = 'b'
	char := 'a'
	fmt.Println("O tipo de char é: ", reflect.TypeOf(char))
	fmt.Println(char)
	fmt.Printf("Tabela ascii= 0x%x\n", char)

}
