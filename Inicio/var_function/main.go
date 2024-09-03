package main

import (
	"fmt"
)

// função principal
func main() {
	//variáveis
	nome, salario := "Marcos", 100
	setName(nome)
	newSalary, bonus := addSalary(salario, 10)
	fmt.Println("Novo salário", newSalary)
	fmt.Println("Bônus", bonus)
}

// funcções
func setName(name string) {
	fmt.Println("Hello", name)
}

// retorno de dois valores
func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

// funcao
func getName() string {
	return "Marcos"
}
