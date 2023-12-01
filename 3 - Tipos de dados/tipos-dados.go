package main

import (
	"errors"
	"fmt"
)

func main() {
	// inteiros: int8, int16, int32, int64
	var numero int = 150
	fmt.Println(numero)

	// reais: float32, float64
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	//string
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// booleano
	var boolean bool = true
	fmt.Println(boolean)

	// tipo erro
	var erro error = errors.New("internal error")
	fmt.Println(erro)
}
