package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// verifica se o serviço está responsivo e tempo de resposta
func main() {
	//tempo
	agora := time.Now()
	//url que será testada por parametro
	url := os.Args[1]
	//acesso a URL
	get, err := http.Get(url)
	//err verificar se deu erro o acesso
	if err != nil {
		fmt.Println("ocorreu um erro ao executar o get url")
		//panic utilizado para crashar a aplicacao, nada mais acontece se der erro
		panic(err)
	}
	//contabiliza o tempo de resposta
	decorrido := time.Since(agora).Seconds()
	//pega o status do get
	status := get.StatusCode
	//imprime as variaveis %d inteiro, %f fload, %s string
	fmt.Printf("Status: [%d] Tempo de carga: [%f]\n", status, decorrido)
}
