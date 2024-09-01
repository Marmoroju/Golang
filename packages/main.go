// cada arquivo criado recebe um nome, por exemplo, ao iniciar um novo arquivo
// com o nome package main
// esse arquivo com outras funções pode ser importado em outros arquivos
// é uma forma de organizar o código

package main

//import "nome do módulo atual/nome do package"
import "pack/funcionarios"

func main() {
	//&nome_do_package.nome_struct
	pessoa := &funcionarios.Pessoa{
		nome:    "Marcos",
		idade:   39,
		salario: 100,
	}
}
