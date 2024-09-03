package main

import "fmt"

type Empregado struct {
	nome string
	id   int
}

func main() {

	emp := Empregado{"Marcos", 1}
	pts := &emp

	fmt.Println(pts)
	pts.nome = "Mourão"
	fmt.Print("Valor pts: ", pts)

}
