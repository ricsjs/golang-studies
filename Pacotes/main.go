package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Esquevendo do arquivo main")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("ricfilho0007@gmail.com")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("E-mail válido")
	}
}
