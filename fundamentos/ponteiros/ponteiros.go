package main

import "fmt"

func main() {
	i := 1

	//Go nao tem aritmetica de ponteiros
	var p_i *int = nil

	p_i = &i //pegando o endereço da variavel
	*p_i++   //desreferenciando (pegando o valor)
	i++

	//Go nao tem aritmetica de ponteiros
	//p++

	fmt.Println(p_i, *p_i, i, &i)
}
