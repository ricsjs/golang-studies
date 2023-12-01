package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel int
	var ponteiro *int

	variavel = 150
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro)  // assim imprime o endereço de memória da variável
	fmt.Println(variavel, *ponteiro) // assim imprime o valor que tá salvo no endereço
}
