package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p := pessoa{"Ricardo", "Alencar", 21, 1.71}
	e := estudante{p, "BSI", "UFRN"}

	fmt.Println(e)
}
