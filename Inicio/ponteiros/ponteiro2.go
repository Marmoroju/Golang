package main

import "fmt"

func pointerFunction(a *int) {
	*a = 400
}

func main() {
	var x = 100
	fmt.Println("O valor da variável x antes da função é: ", x)

	var pa *int = &x
	pointerFunction(pa)
	fmt.Println("O valor de x depois da chamada da função é: ", x)

}
