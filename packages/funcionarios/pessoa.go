// essa package que será importada deverá levar o mesmo nome do diretório
package funcionarios

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func (p *Pessoa) addSalary2(bonus int) {
	p.salario += bonus
}
