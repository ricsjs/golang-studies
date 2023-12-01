package main

import "fmt"

type usuarios struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuarios
	u.nome = "Ricardo"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	u2 := usuarios{"Davi", 21, enderecoExemplo}
	fmt.Println(u2)
}
