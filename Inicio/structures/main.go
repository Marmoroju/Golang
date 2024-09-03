// conceito de ponteiro (mais simples)
// conceito de orientação a objeto na criação de struct
// conceito de crição de tabela de banco de dados na criação de struct

package main

import "fmt"

// DECLARAÇÃO DA STRUCT
type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func main() {

	//OPÇÃO 1 DE DECLARAR STRUCT
	// o & correlaciona o ponteiro definido com o *
	// sem essa correlação e definição é passado apenas a cópia e não o ponteiro

	pessoa := &Pessoa{
		nome:    "Marcos",
		idade:   39,
		salario: 100,
	}

	addSalary(pessoa, 10)
	fmt.Println(pessoa.salario)

	//OPÇÃO 2 DE DECLARAR STRUCT

	pessoa2 := new(Pessoa)
	pessoa2.nome = "Eu"
	pessoa2.idade = 35
	pessoa2.salario = 200

	pessoa2.addSalary2(10)
	fmt.Println(pessoa2.salario)

}

//OPÇÃO 1
// *Pessoa = o * define que será um ponteiro
// o parâmetro p recebe o struct Pessoa como ponteiro e permissão de utilizar seus objetos
func addSalary(p *Pessoa, bonus int) {
	p.salario += bonus
}

//OPÇÃO 2
func (p *Pessoa) addSalary2(bonus int) {
	p.salario += bonus
}
