package main

import "fmt"

func main() {
	canal := make(chan string, 3)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	canal <- "Programando em Go Novamentes!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
