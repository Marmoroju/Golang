//MAPS chave:valor
//Dicionário em python

package main

import "fmt"

func main() {

	//delacaração do maps
	salFuncs := make(map[string]int)
	salFuncs["Marcos"] = 10
	salFuncs["Martha"] = 20

	sal, exists := salFuncs["Marcos"]
	fmt.Println(sal, exists)
}
