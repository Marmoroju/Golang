package main

//worker pool - utilizado quando há uma fila de tarefas a serem utilizadas
// adiciona mais rotinas aquela tarefa para acelerar a execução

import "fmt"

func main() {
	tarefas := make(chan int, 40)
	resultados := make(chan int, 40)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 40; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 40; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

// <- chan (envia dados)  chan <- (recebe dados)  chan (envia e recebe dados)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
