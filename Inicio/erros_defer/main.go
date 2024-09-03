// Go não trabalha com exceptions
// trabalha tratando erro com retorno de valor

package main

import (
	"fmt"
	"os"
)

func main() {

	//defer - fecha arquivos e conexões com bancos
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	/*
		if len(os.Args) != 2 {
			os.Exit(1)
		}
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("erro, não é um número válido!")
			os.Exit(1)
		}
		fmt.Println("número convertido", n)
	*/
}
