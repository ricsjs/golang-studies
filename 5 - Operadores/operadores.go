package main

import "fmt"

func main() {
	// Operadores Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 2 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Operadores de Atribuição
	var variavel1 string = "String"
	variavel := "String 2"

	fmt.Println(variavel, variavel1)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 == 1)
	fmt.Println(2 > 1)
	fmt.Println(2 != 1)
	fmt.Println(2 >= 1)

	// Operadores Lógicos
	fmt.Println("Operadores lógicos")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro && !falso)

	// Operadores Unários
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	// Operador Ternário
	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}

	fmt.Println(texto)
}
