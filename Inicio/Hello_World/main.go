package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args = array
	//ARGS [0,1] = números de argumentos/variáveis
	//Posição 0 do array é o próprio arquivo
	// Posição 1, será passada como argumento no terminal e será concatenada
	// ex.: go run main.go Oi
	// Saída: Hello, Oi
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Hello,", os.Args[1])
}
