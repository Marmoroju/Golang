package main

import "fmt"

func main() {
	var inteiro = 45
	var ponteiro *int = &inteiro
	fmt.Println("Valor da variável inteiro: ", inteiro)
	fmt.Println("Endereço da variável inteiro: ", &inteiro)
	fmt.Println("Valor da variável ponteiro: ", ponteiro)
	fmt.Println("Endereço da variável ponteiro: ", &ponteiro)

	//valor que o ponteiro aponta, que será mais utilizado
	fmt.Println("Valor da variável inteiro pelo ponteiro: ", *ponteiro)

}
