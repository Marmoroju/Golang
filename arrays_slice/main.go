// SLICE = ARRAY DINAMICO

package main

import "fmt"

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func main() {

	// ARRAY
	//var salario [10]int
	//salario[9] = 1000

	// SLICE
	salarios := []int{}

	// adiciona mais 5 posições ao array
	//salarios := make([]int, 5)

	for i := 0; i < 10; i++ {
		//append quando não se sabe o tamanho do slice, array
		salarios = append(salarios, 100+i)
	}
	// o _ serve pra não utilizar o índice ou uma variável
	for _, salario := range salarios {
		fmt.Println(salario)
	}

}
